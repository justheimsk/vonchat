package http

import (
  "fmt"
  "net/http"

  "github.com/justheimsk/vonchat/server/api/v1"
  "github.com/justheimsk/vonchat/server/internal/domain/models"
  "github.com/justheimsk/vonchat/server/internal/infra/config"
  "github.com/justheimsk/vonchat/server/internal/infra/database"
  "github.com/justheimsk/vonchat/server/internal/infra/http/middleware"
  "github.com/go-chi/chi/v5"
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

  loggingMiddleware := middleware.NewLoggingMiddleware(self.logger)
  router := chi.NewRouter()
  if config.Debug {
    router.Use(loggingMiddleware.Run)
  }

  api.LoadHTTPV1Routes(router, self.db, self.logger)
  self.logger.Info("Serving HTTP in port: ", PORT)
  if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", PORT), router); err != nil {
    self.logger.Fatal("Failed to start HTTP server: ", err)
  }
}
