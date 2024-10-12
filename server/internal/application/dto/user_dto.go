package dto

type UserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCreate struct {
	UserDTO
	Password string `json:"password"`
}
