package api

import (
	"github.com/go-chi/chi/v5"
	ws_delivery "github.com/justheimsk/vonchat/server/api/v1/auth/delivery/ws"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	builder "github.com/justheimsk/vonchat/server/internal/builder"
	"github.com/justheimsk/vonchat/server/internal/domain/constants/opcodes"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/http/middleware"
	"github.com/justheimsk/vonchat/server/internal/infra/ws"
)

func LoadHTTPV1Routes(mux *chi.Mux, driver database.DatabaseDriver, logger models.Logger) {
	healthCheckResource := builder.NewHealthBuilder(driver)
	authResource := builder.NewAuthBuilder(driver, logger)
	usersResource := builder.NewUserBuilder(driver, logger)
	authMiddleware := middleware.NewAuthMiddleware(logger, authResource.Service)

	healthCheckResource.Handler.Load(mux)

	mux.Route("/v1", func(router chi.Router) {
		router.Route("/auth", func(authR chi.Router) {
			authResource.Handler.Load(authR)
		})

		router.Route("/", func(protected chi.Router) {
			protected.Use(authMiddleware.Run)

			protected.Route("/users", func(r chi.Router) {
				usersResource.Handler.Load(r)
			})
		})
	})
}

func LoadWSV1Handlers(handler ws.WebsocketHandler, driver database.DatabaseDriver, logger models.Logger) {
	repos := driver.GetRepository()
	authService := service.NewAuthService(repos.Auth, repos.User, logger)
	userService := service.NewUserService(repos.User, logger)

	identifyHandler := ws_delivery.NewIdentifyHandler(authService, userService, logger)
	handler.HandleFunc(opcodes.IDENTIFY, identifyHandler.Handle)
}
