package ws

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

type WebsocketServer struct {
	logger   models.Logger
	upgrader websocket.Upgrader
	Handler  WebsocketHandler
}

func NewWebsocketServer(logger models.Logger) *WebsocketServer {
	return &WebsocketServer{
		logger:  logger,
		Handler: *NewWebsockerHandler(),
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
	self.logger.Info("New Websocket connection")
	for {
		var msg WebsocketMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			self.logger.Error("Failed to read client message", "err", err)
			continue
		}

		writer := NewWebsocketWriter(conn, msg)
		self.Handler.Dispatch(string(writer.Message.T), writer)
	}
}
