package service

import (
	"github.com/justheimsk/vonchat/server/internal/application/dto"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	"github.com/justheimsk/vonchat/server/internal/domain/repository"
)

type UserService struct {
	repo   domain_repo.UserRepository
	logger models.Logger
}

func NewUserService(repo domain_repo.UserRepository, logger models.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (self *UserService) GetUserById(id string) (*dto.UserDTO, error) {
	repo_user, err := self.repo.GetUserById(id)
	if err != nil {
		return nil, models.ErrNotFound
	}

	user := &dto.UserDTO{
		ID:        repo_user.ID,
		Username:  repo_user.Username,
		CreatedAt: repo_user.CreatedAt,
	}

	return user, nil
}

func (self *UserService) GetAll() (*[]dto.UserDTO, error) {
	repo_users, err := self.repo.GetAll()
	if err != nil {
		return nil, models.InternalError
	}

	var users []dto.UserDTO
	for _, user := range *repo_users {
		dto_user := dto.UserDTO{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}

		users = append(users, dto_user)
	}

	return &users, nil
}
