package healthCheckService

import (
	"database/sql"

	healthCheckDelivery "github.com/justheimsk/vonchat/server/api/v1/healthCheck/delivery/http"
	healthCheckTypes "github.com/justheimsk/vonchat/server/api/v1/healthCheck/interfaces"
	controllers "github.com/justheimsk/vonchat/server/internal/controllers/rest"
	repositories "github.com/justheimsk/vonchat/server/internal/repository/pgsql"
)

type healthCheckBuilder struct {
	Handler healthCheckTypes.Handler
}

func New(db *sql.DB) *healthCheckBuilder {
	repo := repositories.NewHealthRepo(db)
	controller := controllers.NewHealthController(repo)

	return &healthCheckBuilder{
		Handler: healthCheckDelivery.NewHTTPHandler(controller),
	}
}
