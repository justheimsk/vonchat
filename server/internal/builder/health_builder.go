package builder

import (
	httpdelivery "github.com/justheimsk/vonchat/server/api/v1/healthCheck/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/persistence/repository"
)

type HealthBuilder struct {
	Handler httpdelivery.HealthHandler
}

func NewHealthBuilder(driver database.DatabaseDriver) *HealthBuilder {
  repo := repository.NewHealthRepository(driver)
	controller := httpdelivery.NewHealthController(repo)

	return &HealthBuilder{
		Handler: *httpdelivery.NewHTTPHandler(controller),
	}
}
