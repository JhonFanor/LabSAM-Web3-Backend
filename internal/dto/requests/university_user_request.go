package requests

type UniversityUserRequest struct {
	Email                 string                `json:"email" validate:"required,email"`
	Password              string                `json:"password" validate:"required,min=8,passwordValidation"`
	Name                  string                `json:"name" validate:"required"`
	UniversityTypeRequest UniversityTypeRequest `json:"university_type" validate:"required"`
	LocationRequest       *LocationRequest      `json:"location"`
	ContactRequest        *ContactRequest       `json:"contact"`
}

type UniversityUserUpdateRequest struct {
	Password              *string                `json:"password"`
	Name                  *string                `json:"name"`
	Avatar                *string                `json:"avatar"`
	UniversityTypeRequest *UniversityTypeRequest `json:"university_type"`
	LocationRequest       *LocationRequest       `json:"location"`
	ContactRequest        *ContactRequest        `json:"contact"`
}
