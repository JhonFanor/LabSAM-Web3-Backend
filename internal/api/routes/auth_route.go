package routes

import (
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, validatorMiddleware *middlewares.ValidatorMiddleware, authController *controllers.AuthController, jwtConfig *config.JwtConfig) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", validatorMiddleware.ValidateInput(&requests.UserRequest{}), authController.RegisterUser)
		auth.POST("/login", validatorMiddleware.ValidateInput(&requests.LoginRequest{}), authController.Login)
		auth.POST("/refresh-token", middlewares.RefreshTokenMiddleware(jwtConfig), authController.RefreshToken)
	}
}
