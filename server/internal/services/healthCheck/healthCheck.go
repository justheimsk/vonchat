package healthCheckService

import (
	"github.com/go-chi/chi/v5"
	healthCheckTypes "github.com/justheimsk/vonchat/server/internal/services/healthCheck/interfaces"
)

type HealthCheckHandler struct {
	controller healthCheckTypes.Controller
}

func New(controller healthCheckTypes.Controller) *HealthCheckHandler {
	return &HealthCheckHandler{
		controller,
	}
}

func (self *HealthCheckHandler) Load(r chi.Router) {
	r.Get("/", self.controller.CheckHealth)
}
