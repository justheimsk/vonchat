package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/justheimsk/vonchat/server/api/v1"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/infra/config"
	"github.com/justheimsk/vonchat/server/internal/infra/database"
	"github.com/justheimsk/vonchat/server/internal/infra/http/middleware"
	"github.com/justheimsk/vonchat/server/internal/infra/ws"
	"github.com/rs/cors"
)

type Server struct {
	db     database.DatabaseDriver
	logger models.Logger
}

func NewServer(db database.DatabaseDriver, logger models.Logger) *Server {
	return &Server{db: db, logger: logger}
}

func (self *Server) Serve(config *config.Config) {
	PORT := config.Port
	self.logger.Infof("Starting HTTP server...")

	loggingMiddleware := middleware.NewLoggingMiddleware(self.logger)
	router := chi.NewRouter()
	if config.Debug {
		router.Use(loggingMiddleware.Run)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	api.LoadHTTPV1Routes(router, self.db, self.logger)

	socket := ws.NewWebsocketServer(self.logger)
	socket.Init(router)

	api.LoadWSV1Handlers(socket.Handler, self.db, self.logger)
	self.logger.Infof("Serving HTTP in port: %s", PORT)

	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", PORT),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		self.logger.Fatal("Failed to start HTTP server", "err", err)
	}
}
