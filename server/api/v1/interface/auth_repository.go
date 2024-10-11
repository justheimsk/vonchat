package interfaces

import "github.com/justheimsk/vonchat/server/api/v1/model"

type AuthRepository interface {
	Register(name string, email string, password string) (model.User, error)
	FetchAccountByEmail(email string) (model.User, error)
	FetchAccountByName(name string) (model.User, error)
}
