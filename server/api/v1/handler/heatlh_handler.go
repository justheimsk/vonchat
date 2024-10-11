package handler

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/interface"
	"github.com/justheimsk/vonchat/server/pkg/concat"
)

type HealthHTTPHandler struct {
	controller interfaces.HealthController
}

func NewHTTPHandler(controller interfaces.HealthController) *HealthHTTPHandler {
	return &HealthHTTPHandler{
		controller,
	}
}

func (self *HealthHTTPHandler) Load(r *http.ServeMux, prefix string) {
	r.HandleFunc(concat.ConcatPath("GET", prefix, ""), self.controller.CheckHealth)
}
