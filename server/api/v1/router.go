package api

import (
	"net/http"

	builder "github.com/justheimsk/vonchat/server/internal/builder"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
)

func LoadV1Routes(r *http.ServeMux, driver database.DatabaseDriver, logger models.Logger) {
	healthCheck := builder.NewHealthBuilder(driver)
	healthCheck.Handler.Load(r, "/v1")

	authResource := builder.NewAuthBuilder(driver, logger)
	authResource.Handler.Load(r, "/v1/auth")
}
