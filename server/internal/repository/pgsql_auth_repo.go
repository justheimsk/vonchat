package repository

import (
	"database/sql"

	"github.com/justheimsk/vonchat/server/api/v1/model"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{
		db,
	}
}

func (self *authRepository) Register(name string, email string, password string) (user model.User, err error) {
	return
}

func (self *authRepository) FetchAccountByName(name string) (user model.User, err error) {
	return
}

func (self *authRepository) FetchAccountByEmail(email string) (user model.User, err error) {
	return
}
