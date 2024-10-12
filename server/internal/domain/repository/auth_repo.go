package domain

import "github.com/justheimsk/vonchat/server/internal/domain/models"

type AuthRepository interface {
	Register(name string, email string, password string) (int, error)
	FetchAccountByEmail(email string) (models.User, error)
}
