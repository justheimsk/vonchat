package domain_service

import "github.com/justheimsk/vonchat/server/internal/application/dto"

type UserService interface {
	GetUserById(string) (*dto.UserDTO, error)
	GetAll() (*[]dto.UserDTO, error)
}
