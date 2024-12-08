package http_delivery

import (
	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	controller AuthController
}

func NewAuthHandler(controller AuthController) *AuthHandler {
	return &AuthHandler{
		controller,
	}
}

func (self *AuthHandler) Load(r chi.Router) {
	r.Post("/register", self.controller.Register)
	r.Post("/login", self.controller.Login)
}
