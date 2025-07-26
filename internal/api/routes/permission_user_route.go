package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PermissionUserRoutesParams struct {
	fx.In
	Router                   *gin.Engine
	TokenMiddleware          *middlewares.TokenMiddleware
	ValidatorMiddleware      *middlewares.ValidatorMiddleware
	PermissionUserController *controllers.PermissionUserController
}

type PermissionUserRoutes struct {
	Router                   *gin.Engine
	TokenMiddleware          *middlewares.TokenMiddleware
	ValidatorMiddleware      *middlewares.ValidatorMiddleware
	PermissionUserController *controllers.PermissionUserController
}

func NewPermissionUserRoutes(p PermissionUserRoutesParams) *PermissionUserRoutes {
	return &PermissionUserRoutes{
		Router:                   p.Router,
		TokenMiddleware:          p.TokenMiddleware,
		ValidatorMiddleware:      p.ValidatorMiddleware,
		PermissionUserController: p.PermissionUserController,
	}
}

func (r *PermissionUserRoutes) Routes() {
	permissionUser := r.Router.Group("/api/permission-user")
	{
		permissionUser.POST("/assign",
			r.TokenMiddleware.ValidateToken(),
			r.ValidatorMiddleware.ValidateInput(&requests.PermissionUserRequest{}),
			r.PermissionUserController.AssignPermissionToUser,
		)

		permissionUser.DELETE("/:permission_id/user/:user_id",
			r.TokenMiddleware.ValidateToken(),
			r.PermissionUserController.RevokePermissionFromUser,
		)

		permissionUser.GET("/user/:user_id",
			r.TokenMiddleware.ValidateToken(),
			r.PermissionUserController.GetAllPermissionsByUser,
		)
	}
}
