package validations

import (
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/pkg/security"

	"github.com/dgrijalva/jwt-go"
)

func ValidateAccessToken(tokenString string, config *config.JwtConfig) (*security.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &security.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*security.Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
