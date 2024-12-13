package ws

import "github.com/gorilla/websocket"

type WebsocketWriter struct {
	Message WebsocketMessage
	Client  *Client
	socket  *websocket.Conn
	server  *WebsocketServer
}

func NewWebsocketWriter(client *Client, socket *websocket.Conn, msg WebsocketMessage, server *WebsocketServer) *WebsocketWriter {
	return &WebsocketWriter{
		Message: msg,
		Client:  client,
		socket:  socket,
		server:  server,
	}
}

func (self *WebsocketWriter) Write(op MSG_OP, t MSG_T, data MSG_DATA) error {
	err := self.socket.WriteJSON(WebsocketMessage{Op: op, T: t, Data: data})
	if err != nil {
		return err
	}

	return nil
}

func (self *WebsocketWriter) GetServer() *WebsocketServer {
	return self.server
}

func (self *WebsocketWriter) CloseSocket() {
	self.server.CloseSocket(self.socket)
}
