package builder

import (
	"database/sql"

	httpdelivery "github.com/justheimsk/vonchat/server/api/v1/healthCheck/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/infra/repository/pgsql"
)

type HealthBuilder struct {
	Handler httpdelivery.HealthHandler
}

func NewHealthBuilder(db *sql.DB) *HealthBuilder {
	repo := pgsql.NewHealthRepository(db)
	controller := httpdelivery.NewHealthController(repo)

	return &HealthBuilder{
		Handler: *httpdelivery.NewHTTPHandler(controller),
	}
}
