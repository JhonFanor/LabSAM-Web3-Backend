package middlewares

import (
	"net/http"
	"strings"

	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/validations"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida el token de acceso en cada solicitud.
func TokenMiddleware(config *config.JwtConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token del encabezado "Authorization"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de acceso no proporcionado"})
			c.Abort()
			return
		}

		// El token debe estar en el formato "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Validar el token
		claims, err := validations.ValidateAccessToken(tokenString, config)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de acceso inválido o expirado"})
			c.Abort()
			return
		}

		// Agregar los claims al contexto para su uso posterior
		c.Set("username", claims.Username)
		c.Next()
	}
}
