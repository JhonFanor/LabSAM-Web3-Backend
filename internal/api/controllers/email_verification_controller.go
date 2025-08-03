package controllers

import (
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type EmailVerificationController struct {
	UserService services.UserService
	JwtConfig   *config.JwtConfig
	Server      *config.ServerConfig
}

func NewEmailVerificationController(us services.UserService, jwt *config.JwtConfig, server *config.ServerConfig) *EmailVerificationController {
	return &EmailVerificationController{
		UserService: us,
		JwtConfig:   jwt,
		Server:      server,
	}
}

func (ctrl *EmailVerificationController) VerifyEmail(c *gin.Context) {
	tokenStr := c.Query("token")
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token no proporcionado"})
		return
	}

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ctrl.JwtConfig.JWT_VERIFY_SECRET), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token inválido o expirado"})
		return
	}

	email := claims.Subject
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email no presente en el token"})
		return
	}

	user, err := ctrl.UserService.FindUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if user.EmailVerified {
		c.JSON(http.StatusOK, gin.H{"message": "Tu correo ya ha sido verificado anteriormente"})
		return
	}

	userVerification := &models.User{
		EmailVerificationToken: "",
		EmailVerified:          true,
		ID:                     user.ID,
	}
	if err := ctrl.UserService.UpdateUser(userVerification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo verificar el correo"})
		return
	}

	c.Redirect(http.StatusFound, ctrl.Server.FRONTEND)

}
