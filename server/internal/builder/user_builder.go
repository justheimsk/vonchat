package builder

import (
	http_delivery "github.com/justheimsk/vonchat/server/api/v1/users/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	domain_service "github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository"
)

type UserBuilder struct {
  Handler http_delivery.UsersHandler
  Controller http_delivery.UsersController
  Repository domain_repo.UserRepository
  Service domain_service.UserService
}

func NewUserBuilder(driver database.DatabaseDriver, logger models.Logger) *UserBuilder {
  repo := repository.NewUserRepository(driver)
  service := service.NewUserService(repo, logger)
  controller := http_delivery.NewUsersController(service)
  handler := http_delivery.NewUsersHandler(*controller)

  return &UserBuilder{
    *handler,
    *controller,
    repo,
    service,
  }
}
