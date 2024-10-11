package controller

import (
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/interface"
)

type AuthController struct {
	authService interfaces.AuthService
}

func NewAuthController(authService interfaces.AuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

func (self *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from auth/resgiter")
}

func (self *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from auth/login")
}
