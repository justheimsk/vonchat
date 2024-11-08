package builder

import (
	httpdelivery "github.com/justheimsk/vonchat/server/api/v1/auth/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository"
)

type AuthBuilder struct {
	Handler httpdelivery.AuthHandler
}

func NewAuthBuilder(driver database.DatabaseDriver, logger models.Logger) *AuthBuilder {
	repo := repository.NewAuthRepository(driver)
	service := service.NewAuthService(repo, logger)
	controller := httpdelivery.NewAuthController(service)

	return &AuthBuilder{
		Handler: *httpdelivery.NewAuthHandler(controller),
	}
}
