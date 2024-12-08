package http_delivery

import (
	"github.com/go-chi/chi/v5"
)

type HealthHandler struct {
	controller HealthController
}

func NewHTTPHandler(controller HealthController) *HealthHandler {
	return &HealthHandler{
		controller,
	}
}

func (self *HealthHandler) Load(r chi.Router) {
	r.Get("/", self.controller.CheckHealth)
}
