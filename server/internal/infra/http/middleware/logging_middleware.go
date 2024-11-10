package middleware

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/infra/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    logger.Log.Debug(r.Method + " " + r.URL.Path)
    next.ServeHTTP(w, r)
  })
}
