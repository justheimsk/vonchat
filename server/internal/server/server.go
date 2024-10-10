package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1"
)

type Server struct {
	db     *sql.DB
	logger *log.Logger
}

func New(db *sql.DB, logger *log.Logger) *Server {
	return &Server{db, logger}
}

func (self *Server) CreateHTTPServer() {
	const PORT int = 8080
	self.logger.Println("Starting HTTP server...")

	router := http.NewServeMux()
	api.LoadV1Routes(router, self.db)

	log.Printf("Serving HTTP in port: %d.\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), router); err != nil {
		log.Fatalf("Failed to start HTTP server: %s", err)
	}
}
