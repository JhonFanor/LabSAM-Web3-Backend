package requests

type BusinessUserRequest struct {
	Email           string           `json:"email" validate:"required,email"`
	Password        string           `json:"password" validate:"required,min=8,passwordValidation"`
	Name            string           `json:"name" validate:"required"`
	Industry        string           `json:"industry"`
	LocationRequest *LocationRequest `json:"location"`
	ContactRequest  *ContactRequest  `json:"contact"`
}

type BusinessUserUpdateRequest struct {
	Password        *string          `json:"password"`
	Name            *string          `json:"name"`
	Avatar          *string          `json:"avatar"`
	Industry        *string          `json:"industry"`
	LocationRequest *LocationRequest `json:"location"`
	ContactRequest  *ContactRequest  `json:"contact"`
}
