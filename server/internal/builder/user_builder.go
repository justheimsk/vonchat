package builder

import (
	"github.com/justheimsk/vonchat/server/api/v1/users/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
)

type UserBuilder struct {
	Handler    http_delivery.UsersHandler
	Controller http_delivery.UsersController
	Repository domain_repo.UserRepository
	Service    domain_service.UserService
}

func NewUserBuilder(driver database.DatabaseDriver, logger models.Logger) *UserBuilder {
	repo := driver.GetRepository().User
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
