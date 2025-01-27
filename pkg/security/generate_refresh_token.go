package security

import (
	"lamsam-web3-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateRefreshToken(username string, config *config.JwtConfig) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 24 horas de expiración

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.REFRESH_SECRET)
}
