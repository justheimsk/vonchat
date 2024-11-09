package builder

import (
	http "github.com/justheimsk/vonchat/server/api/v1/auth/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository"
)

type AuthBuilder struct {
	Handler http.AuthHandler
}

func NewAuthBuilder(driver database.DatabaseDriver, logger models.Logger) *AuthBuilder {
	repo := repository.NewAuthRepository(driver)
	service := service.NewAuthService(repo, logger)
	controller := http.NewAuthController(service)

	return &AuthBuilder{
		Handler: *http.NewAuthHandler(controller),
	}
}
