package middleware

import (
	"context"
	"net/http"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/domain/service"
	"github.com/justheimsk/vonchat/server/pkg/util"
)

type AuthMiddleware struct {
	logger  models.Logger
	service domain_service.AuthService
}

func NewAuthMiddleware(logger models.Logger, service domain_service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		logger:  logger.New("MIDDLEWARE"),
		service: service,
	}
}

func (self *AuthMiddleware) Run(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			util.WriteHTTPError(w, models.ErrUnauthorized)
			return
		}

		token, err := self.service.ValidateToken(header)
		if err != nil {
			util.WriteHTTPError(w, err)
			return
		}

		id, err := self.service.GetIdFromClaims(token)
		if err != nil {
			util.WriteHTTPError(w, err)
			return
		}

		exists := self.service.AccountExists(id)
		if !exists {
			util.WriteHTTPError(w, models.ErrUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
