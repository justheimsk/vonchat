package pgsql

import (
	"database/sql"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db,
	}
}

func (self *AuthRepository) Register(name string, email string, password string) (user models.User, err error) {
	return
}

func (self *AuthRepository) FetchAccountByName(name string) (user models.User, err error) {
	return
}

func (self *AuthRepository) FetchAccountByEmail(email string) (user models.User, err error) {
	return
}
