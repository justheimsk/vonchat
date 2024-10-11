package builder

import (
	"database/sql"

	controller "github.com/justheimsk/vonchat/server/api/v1/controller"
	handler "github.com/justheimsk/vonchat/server/api/v1/handler"
	"github.com/justheimsk/vonchat/server/api/v1/service"
	"github.com/justheimsk/vonchat/server/internal/repository"
)

type AuthBuilder struct {
	Handler handler.AuthHandler
}

func NewAuthBuilder(db *sql.DB) *AuthBuilder {
	repo := repository.NewAuthRepository(db)
	service := service.NewAuthService(repo)
	controller := controller.NewAuthController(service)

	return &AuthBuilder{
		Handler: *handler.NewAuthHandler(controller),
	}
}
