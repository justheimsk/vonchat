package httpdelivery

import (
	"fmt"
	"net/http"

	domain "github.com/justheimsk/vonchat/server/internal/domain/services"
)

type AuthController struct {
	authService domain.AuthService
}

func NewAuthController(authService domain.AuthService) *AuthController {
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
