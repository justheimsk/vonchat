package controller

import (
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/interface"
)

type healthController struct {
	repo interfaces.HealthRepository
}

func NewHealthController(repo interfaces.HealthRepository) *healthController {
	return &healthController{
		repo,
	}
}

func (self *healthController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	ping, err := self.repo.GetPing()
	if err != nil {
		fmt.Fprintf(w, "Failed to get database ping: %s", err)
		return
	}

	fmt.Fprintf(w, "All systems operational, database ping = %dms", ping.Milliseconds())
}
