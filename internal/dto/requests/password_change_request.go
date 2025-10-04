package requests

type PasswordChangeRequest struct {
	Password       string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,passwordValidation"`
}
