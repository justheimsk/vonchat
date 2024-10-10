package healthResource

import (
	"database/sql"

	healthResourceDelivery "github.com/justheimsk/vonchat/server/api/v1/healthCheck/delivery"
	repositories "github.com/justheimsk/vonchat/server/internal/repository/pgsql"
)

type healthCheckBuilder struct {
	Handler healthResourceDelivery.HealthHTTPHandler
}

func New(db *sql.DB) *healthCheckBuilder {
	repo := repositories.NewHealthRepository(db)
	controller := healthResourceDelivery.NewHealthController(repo)

	return &healthCheckBuilder{
		Handler: *healthResourceDelivery.NewHTTPHandler(controller),
	}
}
