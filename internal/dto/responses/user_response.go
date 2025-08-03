package responses

type UserResponse struct {
	ID          uint     `json:"id"`
	Email       string   `json:"email"`
	Avatar      string   `json:"avatar"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

type UserMinimalResponse struct {
	ID             uint                           `json:"id"`
	Avatar         string                         `json:"avatar"`
	RegularUser    *RegularUserMinimalResponse    `json:"regular_user,omitempty"`
	UniversityUser *UniversityUserMinimalResponse `json:"university_user,omitempty"`
	BusinessUser   *BusinessUserMinimalResponse   `json:"business_user,omitempty"`
}

type UserGetAllResponse struct {
	ID             uint                           `json:"id"`
	Email          string                         `json:"email"`
	Avatar         string                         `json:"avatar"`
	RoleID         string                         `json:"role_id"`
	RegularUser    *RegularUserMinimalResponse    `json:"regular_user,omitempty"`
	UniversityUser *UniversityUserMinimalResponse `json:"university_user,omitempty"`
	BusinessUser   *BusinessUserMinimalResponse   `json:"business_user,omitempty"`
}

type UserGetResponse struct {
	ID             uint                       `json:"id"`
	Email          string                     `json:"email"`
	Avatar         string                     `json:"avatar"`
	RoleID         uint                       `json:"role_id"`
	RegularUser    *RegularUserGetResponse    `json:"regular_user,omitempty"`
	UniversityUser *UniversityUserGetResponse `json:"university_user,omitempty"`
	BusinessUser   *BusinessUserGetResponse   `json:"business_user,omitempty"`
}
