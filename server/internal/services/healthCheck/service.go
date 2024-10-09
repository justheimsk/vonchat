package services

import "log"

type healthCheckService struct {
	Handler healthCheckHandler
}

func NewHealthCheckService(logger *log.Logger) *healthCheckService {
	return &healthCheckService{
		Handler: *newHealthCheckHandler(logger, *newHealthCheckController(logger)),
	}
}
