package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/justheimsk/vonchat/server/internal/domain/models"
	domain_repo "github.com/justheimsk/vonchat/server/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
  authRepo   domain_repo.AuthRepository
  userRepo   domain_repo.UserRepository
  logger     models.Logger
}

var secret []byte = []byte("03940943")

func NewAuthService(authRepo domain_repo.AuthRepository, userRepo domain_repo.UserRepository, logger models.Logger) *authService {
  return &authService{
    authRepo:   authRepo,
    userRepo: userRepo,
    logger: logger.New("AUTH SERVICE"),
  }
}

func (self *authService) Register(name string, email string, password string) (token string, err error) {
  _, err = self.authRepo.FetchAccountByEmail(email)
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
    self.logger.Errorf("Failed to hash password: %w", err)
    err = models.InternalError
    return
  }

  id, err := self.authRepo.Register(name, email, pass)
  if err != nil {
    self.logger.Errorf("Failed to register user: %w", err)
    err = models.InternalError
    return
  }

  token, err = self.generateToken(id)
  if err != nil {
    self.logger.Errorf("Failed to generate JWT token: %w", err)
    err = models.InternalError
    return
  }

  self.logger.Debugf("Account create ID=%s", id)
  return
}

func (self *authService) Login(email string, password string) (token string, err error) {
  user, err := self.authRepo.FetchAccountByEmail(email)
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

  self.logger.Debugf("Account login ID=%s", user.ID)
  return
}

func (self *authService) generateToken(id string) (token string, err error) {
  buf := jwt.NewWithClaims(jwt.SigningMethodHS256,
    jwt.MapClaims{
      "id": id,
    })

  token, err = buf.SignedString(secret)
  return
}

func (self *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
  token, err := jwt.Parse(tokenString, func(tk *jwt.Token) (interface{}, error) {
    if _, ok := tk.Method.(*jwt.SigningMethodRSA); ok {
      return "", models.ErrUnauthorized
    }
    return secret, nil
  })

  if err != nil || !token.Valid {
    return nil, models.ErrUnauthorized
  }

  return token, nil
}

func (self *authService) GetIdFromClaims(token *jwt.Token) (string, error) {
  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    self.logger.Errorf("Failed to get ID from claims")
    return "", models.InternalError
  }

  id, ok := claims["id"]
  if !ok || id == nil {
    self.logger.Errorf("Failed to get ID from claims")
    return "", models.InternalError
  }

  return id.(string), nil
}

func (self *authService) hashPassword(password string) (hash string, err error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  if err != nil {
    return
  }

  hash = string(bytes)
  return
}

func (self *authService) AccountExists(id string) bool {
  _, err := self.userRepo.GetUserById(id)
  if err != nil {
    return false
  }

  return true
}
