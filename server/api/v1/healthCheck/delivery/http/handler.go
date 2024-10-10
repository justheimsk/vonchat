package healthCheckDelivery

import (
	"net/http"

	healthCheckTypes "github.com/justheimsk/vonchat/server/api/v1/healthCheck/interfaces"
)

type HealthCheckHTTPHandler struct {
	controller healthCheckTypes.Controller
}

func NewHTTPHandler(controller healthCheckTypes.Controller) *HealthCheckHTTPHandler {
	return &HealthCheckHTTPHandler{
		controller,
	}
}

func (self *HealthCheckHTTPHandler) Load(r *http.ServeMux) {
	r.HandleFunc("GET /", self.controller.CheckHealth)
}
