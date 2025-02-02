package security

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	UserID      uint     `json:"user_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}
