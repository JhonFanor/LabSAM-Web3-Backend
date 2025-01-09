package responses

import "lamsam-web3-backend/internal/models"

type UserResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

type ErrorResponse struct {
	Error   string   `json:"error"`
	Details []string `json:"details,omitempty"`
}
