package sqlite

import (
	"database/sql"

	"github.com/justheimsk/vonchat/server/internal/domain/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db,
	}
}

func (self *userRepository) GetUserById(id string) (*models.User, error) {
	user := &models.User{}

	str := "SELECT * FROM users WHERE id=$1"
	err := self.db.QueryRow(str, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
