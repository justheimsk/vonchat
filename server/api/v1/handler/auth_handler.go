package handler

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/interface"
	"github.com/justheimsk/vonchat/server/pkg/concat"
)

type AuthHandler struct {
	controller interfaces.AuthController
}

func NewAuthHandler(controller interfaces.AuthController) *AuthHandler {
	return &AuthHandler{
		controller,
	}
}

func (self *AuthHandler) Load(r *http.ServeMux, prefix string) {
	r.HandleFunc(concat.ConcatPath("POST", prefix, "/register"), self.controller.Register)
	r.HandleFunc(concat.ConcatPath("POST", prefix, "/login"), self.controller.Login)
}
