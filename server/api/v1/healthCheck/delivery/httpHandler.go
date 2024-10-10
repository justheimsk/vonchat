package healthResourceDelivery

import (
	"net/http"

	healthResourceType "github.com/justheimsk/vonchat/server/api/v1/healthCheck/interfaces"
)

type HealthHTTPHandler struct {
	controller healthResourceType.Controller
}

func NewHTTPHandler(controller healthResourceType.Controller) *HealthHTTPHandler {
	return &HealthHTTPHandler{
		controller,
	}
}

func (self *HealthHTTPHandler) Load(r *http.ServeMux) {
	r.HandleFunc("GET /", self.controller.CheckHealth)
}
