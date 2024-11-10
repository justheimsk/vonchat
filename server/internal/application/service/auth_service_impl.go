package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo   domain.AuthRepository
	logger models.Logger
}

func NewAuthService(repo domain.AuthRepository, logger models.Logger) *authService {
	return &authService{
		repo:   repo,
		logger: logger.New("AUTH SERVICE"),
	}
}

func (self *authService) Register(name string, email string, password string) (token string, err error) {
  start := self.logger.StartTrigger()
	_, err = self.repo.FetchAccountByEmail(email)
	if err == nil {
		err = models.NewCustomError(models.DuplicateErrorCode, "Email already in use.")
		return
	}

  user := models.User{
    Username: name,
    Email: email,
    Password: password,
  }

  if errs := user.Validate(); errs != nil {
    err = models.NewMultiError("400", errs)
    return
  }

	pass, err := self.hashPassword(password)
	if err != nil {
		self.logger.Error(err)
		err = models.InternalError
		return
	}

	id, err := self.repo.Register(name, email, pass)
	if err != nil {
		self.logger.Error(err)
		err = models.InternalError
		return
	}

	token, err = self.generateToken(id)
	if err != nil {
		self.logger.Error(err)
		err = models.InternalError
		return
	}

  self.logger.DebugWithTime(start, "Account create ID=", id)
	return
}

func (self *authService) Login(email string, password string) (token string, err error) {
  start := self.logger.StartTrigger()
	user, err := self.repo.FetchAccountByEmail(email)
	if err != nil {
		err = models.ErrUnauthorized
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		err = models.ErrUnauthorized
		return
	}

	token, err = self.generateToken(user.ID)
	if err != nil {
		err = models.InternalError
		return
	}

  self.logger.DebugWithTime(start, "Account login ID=",user.ID)
	return
}

func (self *authService) generateToken(id int) (token string, err error) {
	buf := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
		})

	token, err = buf.SignedString([]byte("03940943"))
	return
}

func (self *authService) validateToken(token string) (id string, err error) {
	return
}

func (self *authService) hashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return
	}

	hash = string(bytes)
	return
}
