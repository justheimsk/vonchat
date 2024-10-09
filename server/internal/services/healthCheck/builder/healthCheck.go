package builders

import (
	"database/sql"

	controllers "github.com/justheimsk/vonchat/server/internal/controllers/rest"
	repositories "github.com/justheimsk/vonchat/server/internal/repository/pgsql"
	healthCheckService "github.com/justheimsk/vonchat/server/internal/services/healthCheck"
)

type HealthCheckBuilder struct {
	Handler healthCheckService.HealthCheckHandler
}

func NewHealthCheck(db *sql.DB) *HealthCheckBuilder {
	repo := repositories.NewHealthRepo(db)
	controller := controllers.NewHealthController(repo)

	return &HealthCheckBuilder{
		Handler: *healthCheckService.New(controller),
	}
}
