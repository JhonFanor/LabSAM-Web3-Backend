package security

import (
	"lamsam-web3-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateEmailVerificationToken(email string, config *config.JwtConfig) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_VERIFY_SECRET))
}
