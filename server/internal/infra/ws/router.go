package ws

import "github.com/justheimsk/vonchat/server/internal/registry"

type WebsocketRouter struct {
	registry *registry.Registry[string, func(*WebsocketWriter)]
}

func NewWebsockerRouter() *WebsocketRouter {
	return &WebsocketRouter{
		registry: registry.NewRegistry[string, func(*WebsocketWriter)](),
	}
}

func (self *WebsocketRouter) HandleFunc(id string, cb func(*WebsocketWriter)) {
	self.registry.Register(id, cb)
}

func (self *WebsocketRouter) Dispatch(id string, conn *WebsocketWriter) {
	cb, found := self.registry.Get(id)
	if !found {
		return
	}

	cb(conn)
}
