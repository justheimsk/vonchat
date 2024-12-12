package builder

import (
	"github.com/justheimsk/vonchat/server/api/v1/healthCheck/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
)

type HealthBuilder struct {
	Handler    http_delivery.HealthHandler
	Repository domain_repo.HealthRepository
}

func NewHealthBuilder(driver database.DatabaseDriver) *HealthBuilder {
	repo := driver.GetRepository().Health
	controller := http_delivery.NewHealthController(repo)

	return &HealthBuilder{
		Handler:    *http_delivery.NewHTTPHandler(*controller),
		Repository: repo,
	}
}
