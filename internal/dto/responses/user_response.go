package responses

type UserResponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   string `json:"role"`
}

type UserMinimalResponse struct {
	Avatar         string                         `json:"avatar"`
	RegularUser    *RegularUserMinimalResponse    `json:"regular_user,omitempty"`
	UniversityUser *UniversityUserMinimalResponse `json:"university_user,omitempty"`
	BusinessUser   *BusinessUserMinimalResponse   `json:"business_user,omitempty"`
}
