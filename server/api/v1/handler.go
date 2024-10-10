package api

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	healthCheckService "github.com/justheimsk/vonchat/server/api/v1/healthCheck"
)

func LoadV1Routes(r chi.Router, db *sql.DB) {
	healthCheck := healthCheckService.New(db)
	r.Route("/", healthCheck.Handler.Load)
}
