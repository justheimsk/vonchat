package api

import (
	"database/sql"
	"net/http"

	builder "github.com/justheimsk/vonchat/server/internal/builders"
)

func LoadV1Routes(r *http.ServeMux, db *sql.DB) {
	healthCheck := builder.NewHealthBuilder(db)
	healthCheck.Handler.Load(r, "/v1")
}
