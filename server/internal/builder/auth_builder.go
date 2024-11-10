package builder

import (
	delivery_http "github.com/justheimsk/vonchat/server/api/v1/auth/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	domain_service "github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository"
)

type AuthBuilder struct {
	Handler delivery_http.AuthHandler
  Service domain_service.AuthService
  Repository domain_repo.AuthRepository
}

func NewAuthBuilder(driver database.DatabaseDriver, logger models.Logger) *AuthBuilder {
	repo := repository.NewAuthRepository(driver)
	service := service.NewAuthService(repo, logger)
	controller := delivery_http.NewAuthController(service)

	return &AuthBuilder{
		Handler: *delivery_http.NewAuthHandler(controller),
    Service: service,
    Repository: repo,
	}
}
