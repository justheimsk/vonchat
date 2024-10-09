package services

import (
	"log"

	"github.com/go-chi/chi/v5"
)

type healthCheckHandler struct {
	logger     *log.Logger
	controller healthCheckController
}

func newHealthCheckHandler(logger *log.Logger, controller healthCheckController) *healthCheckHandler {
	return &healthCheckHandler{
		logger,
		controller,
	}
}

func (self *healthCheckHandler) Load(r chi.Router) {
	r.Get("/", self.controller.CheckHealth)
	self.logger.Println("HealthCheck handler fully loaded.")
}
