package middlewares

import (
	"net/http"
	"strings"

	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/validations"

	"github.com/gin-gonic/gin"
)

type RefreshTokenMiddleware struct {
	Config *config.JwtConfig
}

func NewRefreshTokenMiddleware(config *config.JwtConfig) *RefreshTokenMiddleware {
	return &RefreshTokenMiddleware{
		Config: config,
	}
}

func (rm *RefreshTokenMiddleware) ValidateRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token no proporcionado"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de refresh token inválido"})
			c.Abort()
			return
		}

		refreshToken := tokenParts[1]

		claims, err := validations.ValidateRefreshToken(refreshToken, rm.Config)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido o expirado"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
