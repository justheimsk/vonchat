package service

import (
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
)

type AuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) *AuthService {
	return &AuthService{
		repo,
	}
}

func (self *AuthService) Register(name string, email string, password string) (user models.User, err error) {
	return
}

func (self *AuthService) FetchAccountByEmail(email string) (user models.User, err error) {
	return
}

func (self *AuthService) FetchAccountByName(name string) (user models.User, err error) {
	return
}

func (self *AuthService) ComparePasswords(password string, hash string) (ok bool, err error) {
	return
}

func (self *AuthService) generateToken(id string) (token string, err error) {
	return
}

func (self *AuthService) validateToken(token string) (id string, err error) {
	return
}

func (self *AuthService) hashPassword(password string) (hash string, err error) {
	return
}
