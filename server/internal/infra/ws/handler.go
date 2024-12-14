package ws

import "github.com/justheimsk/vonchat/server/internal/registry"

type WebsocketHandler struct {
	registry *registry.Registry[int, func(*WebsocketWriter)]
}

func NewWebsockerHandler() *WebsocketHandler {
	return &WebsocketHandler{
		registry: registry.NewRegistry[int, func(*WebsocketWriter)](),
	}
}

func (self *WebsocketHandler) HandleFunc(id int, cb func(*WebsocketWriter)) {
	self.registry.Register(id, cb)
}

func (self *WebsocketHandler) Dispatch(id int, conn *WebsocketWriter) {
	cb, found := self.registry.Get(id)
	if !found {
		return
	}

	cb(conn)
}
