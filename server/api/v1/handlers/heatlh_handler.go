package handler

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/interfaces"
)

type HealthHTTPHandler struct {
	controller interfaces.HealthController
}

func NewHTTPHandler(controller interfaces.HealthController) *HealthHTTPHandler {
	return &HealthHTTPHandler{
		controller,
	}
}

func (self *HealthHTTPHandler) Load(r *http.ServeMux) {
	r.HandleFunc("GET /", self.controller.CheckHealth)
}
