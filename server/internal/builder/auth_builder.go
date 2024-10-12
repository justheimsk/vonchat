package builder

import (
	"database/sql"

	httpdelivery "github.com/justheimsk/vonchat/server/api/v1/auth/delivery/http"
	"github.com/justheimsk/vonchat/server/internal/application/service"
	"github.com/justheimsk/vonchat/server/internal/infra/repository/pgsql"
)

type AuthBuilder struct {
	Handler httpdelivery.AuthHandler
}

func NewAuthBuilder(db *sql.DB) *AuthBuilder {
	repo := pgsql.NewAuthRepository(db)
	service := service.NewAuthService(repo)
	controller := httpdelivery.NewAuthController(service)

	return &AuthBuilder{
		Handler: *httpdelivery.NewAuthHandler(controller),
	}
}
