package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AuthRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	AuthController      *controllers.AuthController
}

type AuthRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	AuthController      *controllers.AuthController
}

func NewAuthRoutes(p AuthRoutesParams) *AuthRoutes {
	return &AuthRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		AuthController:      p.AuthController,
	}
}

func (ar *AuthRoutes) Routes() {
	auth := ar.Router.Group("/api/auth")
	{
		auth.POST("/register/regular", ar.ValidatorMiddleware.ValidateInput(&requests.RegularUserRequest{}), ar.AuthController.RegisterRegularUser)
		auth.POST("/register/university", ar.ValidatorMiddleware.ValidateInput(&requests.UniversityUserRequest{}), ar.AuthController.RegisterUniversityUser)
		auth.POST("/register/business", ar.ValidatorMiddleware.ValidateInput(&requests.BusinessUserRequest{}), ar.AuthController.RegisterBusinessUser)
		auth.POST("/login", ar.ValidatorMiddleware.ValidateInput(&requests.LoginRequest{}), ar.AuthController.Login)
		auth.POST("/token/refresh", ar.AuthController.RefreshToken)
		auth.POST("/logout", ar.AuthController.Logout)
		auth.POST("/forgot-password", ar.AuthController.RequestPasswordReset)
		auth.POST("/reset-password", ar.ValidatorMiddleware.ValidateInput(&requests.ResetPassword{}), ar.AuthController.ResetPassword)
	}
}
