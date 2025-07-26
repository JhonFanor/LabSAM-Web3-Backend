package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UserRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	UserController      *controllers.UserController
}

type UserRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	UserController      *controllers.UserController
}

func NewUserRoutes(p UserRoutesParams) *UserRoutes {
	return &UserRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		UserController:      p.UserController,
	}
}

func (ur *UserRoutes) Routes() {
	user := ur.Router.Group("/api/user")
	{
		user.GET("", ur.TokenMiddleware.ValidateToken(), ur.UserController.GetUserByID)
		user.GET("/admin",
			ur.TokenMiddleware.ValidateToken(),
			ur.UserController.GetAllUsers,
		)

		user.GET("/admin/:id",
			ur.TokenMiddleware.ValidateToken(),
			ur.UserController.GetUserForAdminByID,
		)
	}
}
