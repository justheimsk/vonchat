package server

import (
	"database/sql"
	"log"
)

type Server struct {
  db *sql.DB
  logger *log.Logger
}

func NewServer(db *sql.DB, logger *log.Logger) (*Server) {
  return &Server{ db, logger }
}

func (self *Server) Init() {
  self.logger.Println("Init server");
}
