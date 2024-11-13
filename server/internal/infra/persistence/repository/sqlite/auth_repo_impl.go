package sqlite

import (
  "database/sql"

  "github.com/justheimsk/vonchat/server/internal/domain/models"
)

type authRepository struct {
  db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
  return &authRepository{
    db,
  }
}

func (self *authRepository) Register(name string, email string, password string) (id string, err error) {
  id = ""
  strSql := "INSERT INTO users (username, email, password) values ($1, $2, $3) RETURNING id"
  err = self.db.QueryRow(strSql, name, email, password).Scan(&id)
  return
}

func (self *authRepository) FetchAccountByEmail(email string) (*models.User, error) {
  user := &models.User{}

  strSql := "SELECT id, username, email, password, created_at FROM users WHERE email=$1"
  err := self.db.QueryRow(strSql, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
  if err != nil {
    return nil, err
  }
  return user, nil
}
