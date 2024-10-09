package handlers

import (
	"github.com/go-chi/chi/v5"
	healthCheckService "github.com/justheimsk/vonchat/server/internal/services/healthCheck"
)

type healthCheckHandler struct {
	controller healthCheckService.Controller
}

func NewHealthHandler(controller healthCheckService.Controller) *healthCheckHandler {
	return &healthCheckHandler{
		controller,
	}
}

func (self *healthCheckHandler) Load(r chi.Router) {
	r.Get("/", self.controller.CheckHealth)
}
