package healthCheckDelivery

import (
	"github.com/go-chi/chi/v5"
	healthCheckTypes "github.com/justheimsk/vonchat/server/internal/services/healthCheck/interfaces"
)

type HealthCheckHTTPHandler struct {
	controller healthCheckTypes.Controller
}

func NewHTTPHandler(controller healthCheckTypes.Controller) *HealthCheckHTTPHandler {
	return &HealthCheckHTTPHandler{
		controller,
	}
}

func (self *HealthCheckHTTPHandler) Load(r chi.Router) {
	r.Get("/", self.controller.CheckHealth)
}
