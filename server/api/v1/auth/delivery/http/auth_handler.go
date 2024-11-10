package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/justheimsk/vonchat/server/api/v1/auth"
)

type AuthHandler struct {
	controller auth.Controller
}

func NewAuthHandler(controller auth.Controller) *AuthHandler {
	return &AuthHandler{
		controller,
	}
}

func (self *AuthHandler) Load(r chi.Router) {
	r.Post("/register", self.controller.Register)
	r.Post("/login", self.controller.Login)
}
