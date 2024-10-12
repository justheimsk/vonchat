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

func (self *AuthRepository) Register(name string, email string, password string) (id int, err error) {
	id = 0
	strSql := "INSERT INTO users (username, email, password) values ($1, $2, $3) RETURNING id"

	err = self.db.QueryRow(strSql, name, email, password).Scan(&id)
	return
}

func (self *AuthRepository) FetchAccountByEmail(email string) (user models.User, err error) {
	return
}
