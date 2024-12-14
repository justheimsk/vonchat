package http_delivery

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/domain/service"
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

func (self *UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID")
	if id == "" {
		util.WriteHTTPError(w, models.InternalError)
		return
	}

	users, err := self.service.GetAll()
	if err != nil {
		util.WriteHTTPError(w, err)
		return
	}

	util.WriteHTTPResponse(w, map[string]interface{}{
		"users": users,
	})
}
