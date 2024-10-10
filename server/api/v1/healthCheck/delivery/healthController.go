package healthResourceDelivery

import (
	"fmt"
	"net/http"

	healthTypes "github.com/justheimsk/vonchat/server/api/v1/healthCheck/interfaces"
)

type healthController struct {
	repo healthTypes.Repository
}

func NewHealthController(repo healthTypes.Repository) *healthController {
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
