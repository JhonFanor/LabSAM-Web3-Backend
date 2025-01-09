package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, validatorMiddleware *middlewares.ValidatorMiddleware, authController *controllers.AuthController) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", validatorMiddleware.ValidateInput(&requests.UserRequest{}), authController.RegisterUser)
	}
}
