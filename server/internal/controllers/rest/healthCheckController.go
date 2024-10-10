package restControllers

import (
	"fmt"
	"net/http"

	healthCheckTypes "github.com/justheimsk/vonchat/server/internal/services/healthCheck/interfaces"
)

type healthCheckController struct {
	repo healthCheckTypes.Repository
}

func NewHealthController(repo healthCheckTypes.Repository) *healthCheckController {
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
