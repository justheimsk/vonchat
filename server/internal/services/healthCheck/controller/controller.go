package controllers

import (
	"fmt"
	"net/http"

	healthCheckService "github.com/justheimsk/vonchat/server/internal/services/healthCheck"
)

type healthCheckController struct {
	repo healthCheckService.Repository
}

func NewHealthController(repo healthCheckService.Repository) *healthCheckController {
	return &healthCheckController{
		repo,
	}
}

func (self *healthCheckController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	ping, err := self.repo.GetPing()
	if err != nil {
		fmt.Fprintf(w, "Failed to get database ping: %s", err)
		return
	}

	fmt.Fprintf(w, "All systems operational, database ping = %dms", ping.Milliseconds())
}
