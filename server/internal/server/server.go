package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	services "github.com/justheimsk/vonchat/server/internal/services/healthCheck"
)

type Server struct {
  db *sql.DB
  logger *log.Logger
}

func NewServer(db *sql.DB, logger *log.Logger) (*Server) {
  return &Server{ db, logger }
}

func (self *Server) Init() {
  self.logger.Println("Starting HTTP server...")
  
  mainRouter := chi.NewRouter()
  healthCheckService := services.NewHealthCheckService(self.logger)
  mainRouter.Route("/", healthCheckService.Handler.Load)

  if err := http.ListenAndServe("0.0.0.0:8080", mainRouter); err != nil {
    log.Fatalf("Failed to start HTTP server: %s", err)
  }
}
