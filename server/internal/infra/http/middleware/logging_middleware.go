package middleware

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

type LoggingMiddleware struct {
	logger models.Logger
}

func NewLoggingMiddleware(logger models.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

func (self *LoggingMiddleware) Run(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		self.logger.Debugf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
