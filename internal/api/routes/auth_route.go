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
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	RefreshTokenMiddleware *middlewares.RefreshTokenMiddleware
	AuthController         *controllers.AuthController
}

type AuthRoutes struct {
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	RefreshTokenMiddleware *middlewares.RefreshTokenMiddleware
	AuthController         *controllers.AuthController
}

func NewAuthRoutes(p AuthRoutesParams) *AuthRoutes {
	return &AuthRoutes{
		Router:                 p.Router,
		ValidatorMiddleware:    p.ValidatorMiddleware,
		RefreshTokenMiddleware: p.RefreshTokenMiddleware,
		AuthController:         p.AuthController,
	}
}

func (ar *AuthRoutes) Routes() {
	auth := ar.Router.Group("/auth")
	{
		auth.POST("/regular/register", ar.ValidatorMiddleware.ValidateInput(&requests.RegularUserRequest{}), ar.AuthController.RegisterRegularUser)
		auth.POST("/university/register", ar.ValidatorMiddleware.ValidateInput(&requests.UniversityUserRequest{}), ar.AuthController.RegisterUniversityUser)
		auth.POST("/business/register", ar.ValidatorMiddleware.ValidateInput(&requests.BusinessUserRequest{}), ar.AuthController.RegisterBusinessUser)
		auth.POST("/login", ar.ValidatorMiddleware.ValidateInput(&requests.LoginRequest{}), ar.AuthController.Login)
		auth.POST("token/refresh", ar.RefreshTokenMiddleware.ValidateRefreshToken(), ar.AuthController.RefreshToken)
	}
}
