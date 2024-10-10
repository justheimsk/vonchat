package interfaces

import "github.com/justheimsk/vonchat/server/api/v1/models"

type AuthRepository interface {
	Register(name string, email string, password string) (models.User, error)
	FetchAccountByEmail(email string) (models.User, error)
	FetchAccountByName(name string) (models.User, error)
}
