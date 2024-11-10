package delivery_http

import (
	"fmt"
	"net/http"

	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
)

type healthController struct {
	repo domain.HealthRepository
}

func NewHealthController(repo domain.HealthRepository) *healthController {
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
