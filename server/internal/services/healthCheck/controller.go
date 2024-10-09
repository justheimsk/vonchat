package healthCheckService

import (
	"fmt"
	"log"
	"net/http"
)

type healthCheckController struct {
	logger *log.Logger
	repo   *healthCheckRepo
}

func NewController(logger *log.Logger, repo *healthCheckRepo) *healthCheckController {
	return &healthCheckController{
		logger,
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
