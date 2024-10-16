package httpdelivery

import (
	"encoding/json"
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/application/dto"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain "github.com/justheimsk/vonchat/server/internal/domain/services"
	"github.com/justheimsk/vonchat/server/pkg/util"
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
	var account dto.UserCreate

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		util.WriteHTTPError(w, models.InternalError)
		return
	}

	token, err := self.authService.Register(account.Username, account.Email, account.Password)
	if err != nil {
		util.WriteHTTPError(w, err)
		return
	}

	util.WriteHTTPResponse(w, map[string]interface{}{
		"token": token,
	})
}

func (self *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var account dto.UserCreate

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		util.WriteHTTPError(w, models.InternalError)
		return
	}

	token, err := self.authService.Login(account.Email, account.Password)
	if err != nil {
		util.WriteHTTPError(w, err)
		return
	}

	util.WriteHTTPResponse(w, map[string]interface{}{
		"token": token,
	})
}
