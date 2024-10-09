package healthCheckService

import (
	"fmt"
	"log"
	"net/http"
)

type healthCheckController struct {
	logger *log.Logger
}

func NewController(logger *log.Logger) *healthCheckController {
	return &healthCheckController{
		logger,
	}
}

func (self *healthCheckController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All systems operational.")
}
