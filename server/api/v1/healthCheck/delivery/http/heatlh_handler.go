package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/justheimsk/vonchat/server/api/v1/healthCheck"
)

type HealthHandler struct {
	controller healthCheck.Controller
}

func NewHTTPHandler(controller healthCheck.Controller) *HealthHandler {
	return &HealthHandler{
		controller,
	}
}

func (self *HealthHandler) Load(r chi.Router) {
	r.Get("/", self.controller.CheckHealth)
}
