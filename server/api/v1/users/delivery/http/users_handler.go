package http_delivery

import "github.com/go-chi/chi/v5"

type UsersHandler struct {
	controller UsersController
}

func NewUsersHandler(controller UsersController) *UsersHandler {
	return &UsersHandler{
		controller,
	}
}

func (self *UsersHandler) Load(r chi.Router) {
	r.Get("/@me", self.controller.GetMe)
}
