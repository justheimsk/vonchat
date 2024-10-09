package healthCheckService

import (
	"database/sql"

	controllers "github.com/justheimsk/vonchat/server/internal/controllers/rest"
	repositories "github.com/justheimsk/vonchat/server/internal/repository/pgsql"
)

type HealthCheckBuilder struct {
	Handler HealthCheckHandler
}

func New(db *sql.DB) *HealthCheckBuilder {
	repo := repositories.NewHealthRepo(db)
	controller := controllers.NewHealthController(repo)

	return &HealthCheckBuilder{
		Handler: *NewHandler(controller),
	}
}
