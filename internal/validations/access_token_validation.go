package validations

import (
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/pkg/security"

	"github.com/dgrijalva/jwt-go"
)

func ValidateAccessToken(tokenString string, config *config.JwtConfig) (*security.Claims, error) {
	claims := &security.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
