package http

import (
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/http/middleware"
)

type Server struct {
	db     database.DatabaseDriver
	logger models.Logger
}

func NewServer(db database.DatabaseDriver, logger models.Logger) *Server {
	return &Server{db: db, logger: logger.New("HTTP")}
}

func (self *Server) Serve(config *config.Config) {
  PORT := config.Port
	self.logger.Info("Starting HTTP server...")

	router := http.NewServeMux()
	api.LoadV1Routes(router, self.db, self.logger)

  var mux http.Handler

  if config.Debug {
    mux = use(router, middleware.LoggingMiddleware)
  } else {
    mux = router
  }
    
	self.logger.Info("Serving HTTP in port: ", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", PORT), mux); err != nil {
		self.logger.Fatal("Failed to start HTTP server: ", err)
	}
}

func use(r *http.ServeMux, middlewares ...func(next http.Handler) http.Handler) http.Handler {
	var s http.Handler
	s = r

	for _, mw := range middlewares {
		s = mw(s)
	}

	return s
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
