package interfaces

import (
	"net/http"

	"github.com/justheimsk/vonchat/server/api/v1/model"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type AuthRepository interface {
	Register(name string, email string, password string) (model.User, error)
	FetchAccountByEmail(email string) (model.User, error)
	FetchAccountByName(name string) (model.User, error)
}

type AuthService interface {
	AuthRepository
	ComparePasswords(password string, hash string) (bool, error)
}
