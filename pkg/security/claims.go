package security

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID      uint     `json:"user_id"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}
