package requests

type UniversityUserRequest struct {
	Email           string                 `json:"email" validate:"required,email"`
	Password        string                 `json:"password" validate:"required,min=8,passwordValidation"`
	Name            string                 `json:"name" validate:"required"`
	UniversityType  *UniversityTypeRequest `json:"university_type"`
	LocationRequest *LocationRequest       `json:"location"`
	ContactRequest  *ContactRequest        `json:"contact"`
}
