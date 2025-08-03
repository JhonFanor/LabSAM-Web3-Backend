package security

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenStr string, secret string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token inválido o expirado")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("no se pudieron obtener los claims")
	}

	return claims, nil
}
