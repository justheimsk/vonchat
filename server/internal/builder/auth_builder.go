package builder

import (
	http_delivery "github.com/justheimsk/vonchat/server/api/v1/auth/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
)

type AuthBuilder struct {
	Handler        http_delivery.AuthHandler
	Service        domain_service.AuthService
	AuthRepository domain_repo.AuthRepository
	UserRepository domain_repo.UserRepository
}

func NewAuthBuilder(driver database.DatabaseDriver, logger models.Logger) *AuthBuilder {
	repos := driver.GetRepository()
	service := service.NewAuthService(repos.Auth, repos.User, logger)
	controller := http_delivery.NewAuthController(service)

	return &AuthBuilder{
		Handler:        *http_delivery.NewAuthHandler(*controller),
		Service:        service,
		AuthRepository: repos.Auth,
		UserRepository: repos.User,
	}
}
