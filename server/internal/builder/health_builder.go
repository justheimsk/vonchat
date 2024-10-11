package builder

import (
	"database/sql"

	controller "github.com/justheimsk/vonchat/server/api/v1/controller"
	handler "github.com/justheimsk/vonchat/server/api/v1/handler"
	"github.com/justheimsk/vonchat/server/internal/repository"
)

type HealthBuilder struct {
	Handler handler.HealthHTTPHandler
}

func NewHealthBuilder(db *sql.DB) *HealthBuilder {
	repo := repository.NewHealthRepository(db)
	controller := controller.NewHealthController(repo)

	return &HealthBuilder{
		Handler: *handler.NewHTTPHandler(controller),
	}
}
