package healthCheckService

import (
	"database/sql"

	controllers "github.com/justheimsk/vonchat/server/internal/controllers/rest"
	repositories "github.com/justheimsk/vonchat/server/internal/repository/pgsql"
	healthCheckDelivery "github.com/justheimsk/vonchat/server/internal/services/healthCheck/delivery/http"
)

type healthCheckBuilder struct {
	Handler healthCheckDelivery.HealthCheckHTTPHandler
}

func New(db *sql.DB) *healthCheckBuilder {
	repo := repositories.NewHealthRepo(db)
	controller := controllers.NewHealthController(repo)

	return &healthCheckBuilder{
		Handler: *healthCheckDelivery.NewHTTPHandler(controller),
	}
}
