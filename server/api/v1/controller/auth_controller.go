package controller

import (
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/interface"
)

type AuthController struct {
	repo interfaces.AuthRepository
}

func NewAuthController(repo interfaces.AuthRepository) *AuthController {
	return &AuthController{
		repo,
	}
}

func (self *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from auth/resgiter")
}

func (self *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from auth/login")
}
