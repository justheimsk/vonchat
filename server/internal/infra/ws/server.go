package ws

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/justheimsk/vonchat/server/internal/domain/constants/opcodes"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/registry"
)

type WebsocketServer struct {
	logger   models.Logger
	upgrader websocket.Upgrader
	Handler  WebsocketHandler
	Sockets  *registry.Registry[string, *Client]
}

func NewWebsocketServer(logger models.Logger) *WebsocketServer {
	return &WebsocketServer{
		logger:  logger,
		Handler: *NewWebsockerHandler(),
		Sockets: registry.NewRegistry[string, *Client](),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
	}
}

func (self *WebsocketServer) Init(router chi.Router) {
	router.Get("/ws", self.Upgrade)
}

func (self *WebsocketServer) Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := self.upgrader.Upgrade(w, r, nil)
	if err != nil {
		self.logger.Error("Failed to upgrade request to Websocket", "err", err)
		return
	}

	defer conn.Close()
	self.logger.Infof("Client %s connected", conn.RemoteAddr().String())

	client := NewClient(conn.RemoteAddr().String(), conn)
	err = self.Sockets.Register(client.RandomID, client)
	if err != nil {
		self.logger.Error("Failed to add socket to registry", "err", err)
		return
	}

	time.AfterFunc(60*time.Second, func() {
		if _, found := self.Sockets.Get(conn.RemoteAddr().String()); found {
			if !client.Authenticated {
				self.logger.Infof("Client %s took too long to identify, closing connection...", conn.RemoteAddr().String())
				self.CloseSocket(conn)
			}
		}
	})

	for {
		var msg WebsocketMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
				self.logger.Error("Failed to read message:", "err", err)
			}

			self.CloseSocket(conn)
			break
		}

		writer := NewWebsocketWriter(client, conn, msg, self)
		if writer.Message.Op != opcodes.IDENTIFY && !client.Authenticated {
			self.logger.Infof("Client %s unauthorized", conn.RemoteAddr().String())
			self.CloseSocket(conn)
		}

		self.Handler.Dispatch(int(writer.Message.Op), writer)
	}
}

func (self *WebsocketServer) Broadcast(msg *WebsocketMessage) {
	for _, client := range self.Sockets.Values() {
		if !client.Authenticated {
			continue
		}

		err := client.Socket.WriteJSON(msg)
		if err != nil {
			self.logger.Errorf("Failed to broadcast message to %s(%s)", client.user.Username, client.RandomID)
		}
	}
}

func (self *WebsocketServer) CloseSocket(socket *websocket.Conn) {
	socket.Close()

	if _, found := self.Sockets.Get(socket.RemoteAddr().String()); found {
		self.Sockets.Remove(socket.RemoteAddr().String())
		self.logger.Infof("Client %s closed connection", socket.RemoteAddr().String())
	}
}
