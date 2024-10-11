package domain

import domain "github.com/justheimsk/vonchat/server/internal/domain/repository"

type AuthService interface {
	domain.AuthRepository
	ComparePasswords(password string, hash string) (bool, error)
}
