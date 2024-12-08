package http_delivery

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_service "github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/pkg/util"
)

type UsersController struct {
	service domain_service.UserService
}

func NewUsersController(service domain_service.UserService) *UsersController {
	return &UsersController{
		service,
	}
}

func (self *UsersController) GetMe(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID")
	if id == "" {
		util.WriteHTTPError(w, models.InternalError)
		return
	}

	user, err := self.service.GetUserById(id.(string))
	if err != nil {
		util.WriteHTTPError(w, models.InternalError)
		return
	}

	util.WriteHTTPResponse(w, map[string]interface{}{
		"user": user,
	})
}
