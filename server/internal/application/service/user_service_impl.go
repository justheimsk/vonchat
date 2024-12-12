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
		logger: logger.New("USER SERVICE"),
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
		Email:     repo_user.Email,
		CreatedAt: repo_user.CreatedAt,
	}

	return user, nil
}
