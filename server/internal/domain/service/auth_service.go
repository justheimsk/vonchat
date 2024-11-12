package domain_service

import "github.com/golang-jwt/jwt/v5"

type AuthService interface {
	Register(name string, email string, password string) (string, error)
	Login(email string, password string) (string, error)
  ValidateToken(string) (*jwt.Token, error)
  GetIdFromClaims(*jwt.Token) (string, error)
}
