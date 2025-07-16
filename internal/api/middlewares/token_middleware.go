package middlewares

import (
	"net/http"
	"strings"

	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/validations"

	"github.com/gin-gonic/gin"
)

type TokenMiddleware struct {
	Config *config.JwtConfig
}

func NewTokenMiddleware(config *config.JwtConfig) *TokenMiddleware {
	return &TokenMiddleware{
		Config: config,
	}
}

func (tm *TokenMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de acceso no proporcionado"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		claims, err := validations.ValidateAccessToken(tokenString, tm.Config)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de acceso inválido o expirado"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func (tm *TokenMiddleware) ValidateOptionalToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Next()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Next()
			return
		}

		tokenString := tokenParts[1]

		claims, err := validations.ValidateAccessToken(tokenString, tm.Config)
		if err != nil {
			c.Next()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
