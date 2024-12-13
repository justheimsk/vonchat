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
		clientIp := r.Header.Get("X-FORWARDED-FOR")
		if clientIp == "" {
			clientIp = r.RemoteAddr
		}

		self.logger.Debugf("NEW REQUEST method=%s path=%s from=%s", r.Method, r.URL.Path, clientIp)
		next.ServeHTTP(w, r)
	})
}
