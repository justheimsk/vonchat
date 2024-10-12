package api

import (
	"database/sql"
	"net/http"

	builder "github.com/justheimsk/vonchat/server/internal/builder"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

func LoadV1Routes(r *http.ServeMux, db *sql.DB, logger models.Logger) {
	healthCheck := builder.NewHealthBuilder(db)
	healthCheck.Handler.Load(r, "/v1")

	authResource := builder.NewAuthBuilder(db, logger)
	authResource.Handler.Load(r, "/v1/auth")
}
