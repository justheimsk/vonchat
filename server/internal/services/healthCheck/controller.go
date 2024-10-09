package services

import (
	"fmt"
	"log"
	"net/http"
)

type healthCheckController struct {
	logger *log.Logger
}

func newHealthCheckController(logger *log.Logger) *healthCheckController {
	return &healthCheckController{
		logger,
	}
}

func (self *healthCheckController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	self.logger.Println("Controller / requested")
	fmt.Fprintf(w, "Hello World!")
}
