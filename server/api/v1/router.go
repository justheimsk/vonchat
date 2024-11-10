package api

import (
	"github.com/go-chi/chi/v5"
	builder "github.com/justheimsk/vonchat/server/internal/builder"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
)

func LoadV1Routes(mux *chi.Mux, driver database.DatabaseDriver, logger models.Logger) {
	healthCheckResource := builder.NewHealthBuilder(driver)
	authResource := builder.NewAuthBuilder(driver, logger)

  mux.Route("/v1", func(router chi.Router) {
	  healthCheckResource.Handler.Load(router)

    mux.Route("/v1/auth", func(authR chi.Router) {
	    authResource.Handler.Load(authR)
    })
  })
}
