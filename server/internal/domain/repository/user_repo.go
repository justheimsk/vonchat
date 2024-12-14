package domain_repo

import "github.com/justheimsk/vonchat/server/internal/domain/models"

type UserRepository interface {
	GetUserById(string) (*models.User, error)
	GetAll() (*[]models.User, error)
}
