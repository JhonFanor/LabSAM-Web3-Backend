package requests

type BusinessUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,passwordValidation"`
	Name     string `json:"name" validate:"required"`
}
