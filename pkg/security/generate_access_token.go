package security

import (
	"lamsam-web3-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(username string, config *config.JwtConfig) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute) // 15 minutos de expiración

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRET_KEY)
}
