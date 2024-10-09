package healthCheckService

import (
	"database/sql"
	"log"
)

type healthCheckService struct {
	logger  *log.Logger
	Handler *healthCheckHandler
}

func New(logger *log.Logger, db *sql.DB) *healthCheckService {
	repo := NewRepo(db)
	controller := NewController(logger, repo)

	return &healthCheckService{
		logger:  logger,
		Handler: NewHandler(logger, *controller),
	}
}
