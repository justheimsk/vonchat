package ws

import "github.com/gorilla/websocket"

type WebsocketWriter struct {
	Message WebsocketMessage
	conn    *websocket.Conn
}

func NewWebsocketWriter(conn *websocket.Conn, msg WebsocketMessage) *WebsocketWriter {
	return &WebsocketWriter{
		Message: msg,
		conn:    conn,
	}
}

func (self *WebsocketWriter) Write(op MSG_OP, t MSG_T, data MSG_DATA) error {
	msg := NewWebsocketMessage(op, t, data)
	err := self.conn.WriteJSON(msg)
	if err != nil {
		return err
	}

	return nil
}
