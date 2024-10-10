package api

import (
	"database/sql"
	"net/http"

	healthCheckService "github.com/justheimsk/vonchat/server/api/v1/healthCheck"
)

func LoadV1Routes(r *http.ServeMux, db *sql.DB) {
	healthCheck := healthCheckService.New(db)
	healthCheck.Handler.Load(r)
}
