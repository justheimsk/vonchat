package ws

import "github.com/justheimsk/vonchat/server/internal/registry"

type WebsocketHandler struct {
	registry *registry.Registry[string, func(*WebsocketWriter)]
}

func NewWebsockerHandler() *WebsocketHandler {
	return &WebsocketHandler{
		registry: registry.NewRegistry[string, func(*WebsocketWriter)](),
	}
}

func (self *WebsocketHandler) HandleFunc(id string, cb func(*WebsocketWriter)) {
	self.registry.Register(id, cb)
}

func (self *WebsocketHandler) Dispatch(id string, conn *WebsocketWriter) {
	cb, found := self.registry.Get(id)
	if !found {
		return
	}

	cb(conn)
}
