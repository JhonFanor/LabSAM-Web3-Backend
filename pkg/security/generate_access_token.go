package security

import (
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(userID uint, username string, email string, role models.Role, permissions []models.Permission, config *config.JwtConfig) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	var permissionNames []string
	for _, p := range permissions {
		permissionNames = append(permissionNames, p.Name)
	}

	claims := &Claims{
		UserID:      userID,
		Username:    username,
		Email:       email,
		Role:        role.Name,
		Permissions: permissionNames,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRET_KEY)
}
