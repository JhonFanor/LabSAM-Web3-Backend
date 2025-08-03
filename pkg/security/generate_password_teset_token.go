package security

import (
	"lamsam-web3-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GeneratePasswordResetToken(email string, jwtConfig *config.JwtConfig) (string, error) {
	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.SECRET_KEY))
}
