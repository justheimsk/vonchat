package ws_delivery

import "github.com/justheimsk/vonchat/server/internal/infra/ws"

type IdentifyHandler struct{}

func NewIdentifyHandler() *IdentifyHandler {
	return &IdentifyHandler{}
}

func (self *IdentifyHandler) Handle(w *ws.WebsocketWriter) {
	w.Write(0, "IDENTIFY_OK", nil)
}
