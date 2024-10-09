package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/justheimsk/vonchat/server/internal/services/healthCheck"
)

type Server struct {
	db     *sql.DB
	logger *log.Logger
}

func NewServer(db *sql.DB, logger *log.Logger) *Server {
	return &Server{db, logger}
}

func (self *Server) Init() {
	const PORT int = 8080

	self.logger.Println("Starting HTTP server...")
	mainRouter := chi.NewRouter()

	healthCheck := healthCheckService.NewHandler(self.logger, *healthCheckService.NewController(self.logger))
	mainRouter.Route("/", healthCheck.Load)

	log.Printf("Serving HTTP in port: %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), mainRouter); err != nil {
		log.Fatalf("Failed to start HTTP server: %s", err)
	}
}
