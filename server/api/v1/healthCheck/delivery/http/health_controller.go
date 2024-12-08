package http_delivery

import (
	"fmt"
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"github.com/justheimsk/vonchat/server/pkg/util"
)

type HealthController struct {
  repo domain.HealthRepository
}

func NewHealthController(repo domain.HealthRepository) *HealthController {
  return &HealthController{
    repo,
  }
}

func (self *HealthController) CheckHealth(w http.ResponseWriter, _ *http.Request) {
  ping, err := self.repo.GetPing()
  if err != nil {
    util.WriteHTTPError(w, models.InternalError)
    return
  }

  msg := fmt.Sprintf("All systems operational, database ping = %dms", ping.Milliseconds())
  util.WriteHTTPResponse(w, map[string]interface{}{
    "message": msg,
    "version": "1",
  })
}
