package dto

type UserDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

type UserCreate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
