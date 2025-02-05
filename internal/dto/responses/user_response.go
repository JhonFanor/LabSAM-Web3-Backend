package responses

import "lamsam-web3-backend/internal/models"

type UserResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}
