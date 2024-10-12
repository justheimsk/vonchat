package service

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) *AuthService {
	return &AuthService{
		repo,
	}
}

func (self *AuthService) Register(name string, email string, password string) (token string, err error) {
	_, err = self.repo.FetchAccountByEmail(email)
	if err == nil {
		err = models.NewCustomError(models.DuplicateErrorCode, "Email already in use.")
		return
	}

	pass, err := self.hashPassword(password)
	if err != nil {
		err = models.InternalError
		return
	}

	id, err := self.repo.Register(name, email, pass)
	if err != nil {
		err = models.InternalError
		return
	}

	token, err = self.generateToken(id)
	return
}

func (self *AuthService) Login(email string, password string) (token string, err error) {
	return
}

func (self *AuthService) generateToken(id int) (token string, err error) {
	buf := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
		})

	token, err = buf.SignedString([]byte("03940943"))
	return
}

func (self *AuthService) validateToken(token string) (id string, err error) {
	return
}

func (self *AuthService) hashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	hash = string(bytes)

	return
}
