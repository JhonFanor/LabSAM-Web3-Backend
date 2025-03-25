package responses

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}

type UserMinimalResponse struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
