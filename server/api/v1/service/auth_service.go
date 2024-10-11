package service

import (
	interfaces "github.com/justheimsk/vonchat/server/api/v1/interface"
	"github.com/justheimsk/vonchat/server/api/v1/model"
)

type AuthService struct {
	repo interfaces.AuthRepository
}

func NewAuthService(repo interfaces.AuthRepository) *AuthService {
	return &AuthService{
		repo,
	}
}

func (self *AuthService) Register(name string, email string, password string) (user model.User, err error) {
	return
}

func (self *AuthService) FetchAccountByEmail(email string) (user model.User, err error) {
	return
}

func (self *AuthService) FetchAccountByName(name string) (user model.User, err error) {
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
