package builder

import (
	http_delivery "github.com/justheimsk/vonchat/server/api/v1/healthCheck/delivery/http"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository"
)

type HealthBuilder struct {
	Handler    http_delivery.HealthHandler
	Repository domain_repo.HealthRepository
}

func NewHealthBuilder(driver database.DatabaseDriver) *HealthBuilder {
	repo := repository.NewHealthRepository(driver)
	controller := http_delivery.NewHealthController(repo)

	return &HealthBuilder{
		Handler:    *http_delivery.NewHTTPHandler(*controller),
		Repository: repo,
	}
}
