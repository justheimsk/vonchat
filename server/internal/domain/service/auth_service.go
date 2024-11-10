package domain_service

type AuthService interface {
	Register(name string, email string, password string) (string, error)
	Login(email string, password string) (string, error)
}
