package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

type Server struct {
	db *sql.DB
  logger models.Logger
}

func New(db *sql.DB, logger models.Logger) *Server {
  return &Server{db: db, logger: logger.New("HTTP")}
}

func (self *Server) CreateHTTPServer() {
	const PORT int = 8080

	self.logger.Info("Starting HTTP server...")

	router := http.NewServeMux()
	api.LoadV1Routes(router, self.db, self.logger)

	self.logger.Info("Serving HTTP in port: ", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), router); err != nil {
		self.logger.Fatal("Failed to start HTTP server: ", err)
	}
}

// func ensureTrailingSlash(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path != "/" && endsWithSlash(r.URL.Path) {
// 			path := r.URL.Path[:len(r.URL.Path)-1]
// 			http.Redirect(w, r, path, http.StatusMovedPermanently)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
//
// func endsWithSlash(path string) bool {
// 	return len(path) > 0 && path[len(path)-1] == '/'
// }
