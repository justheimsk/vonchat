package http

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/auth"
	"github.com/justheimsk/vonchat/server/pkg/util"
)

type AuthHandler struct {
	controller auth.Controller
}

func NewAuthHandler(controller auth.Controller) *AuthHandler {
	return &AuthHandler{
		controller,
	}
}

func (self *AuthHandler) Load(r *http.ServeMux, prefix string) {
	r.HandleFunc(util.ConcatPath("POST", prefix, "/register"), self.controller.Register)
	r.HandleFunc(util.ConcatPath("POST", prefix, "/login"), self.controller.Login)
}
