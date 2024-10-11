package httpdelivery

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/healthCheck"
	"github.com/justheimsk/vonchat/server/pkg/concat"
)

type HealthHandler struct {
	controller healthCheck.Controller
}

func NewHTTPHandler(controller healthCheck.Controller) *HealthHandler {
	return &HealthHandler{
		controller,
	}
}

func (self *HealthHandler) Load(r *http.ServeMux, prefix string) {
	r.HandleFunc(concat.ConcatPath("GET", prefix, ""), self.controller.CheckHealth)
}
