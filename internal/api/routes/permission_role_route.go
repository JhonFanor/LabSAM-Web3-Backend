package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PermissionRoleRoutesParams struct {
	fx.In
	Router                   *gin.Engine
	TokenMiddleware          *middlewares.TokenMiddleware
	PermissionRoleController *controllers.PermissionRoleController
}

type PermissionRoleRoutes struct {
	Router                   *gin.Engine
	TokenMiddleware          *middlewares.TokenMiddleware
	PermissionRoleController *controllers.PermissionRoleController
}

func NewPermissionRoleRoutes(p PermissionRoleRoutesParams) *PermissionRoleRoutes {
	return &PermissionRoleRoutes{
		Router:                   p.Router,
		TokenMiddleware:          p.TokenMiddleware,
		PermissionRoleController: p.PermissionRoleController,
	}
}

func (r *PermissionRoleRoutes) Routes() {
	permissionRole := r.Router.Group("/api/permission-role")
	{
		permissionRole.GET("/role/:role_id/user/:user_id",
			r.TokenMiddleware.ValidateToken(),
			r.PermissionRoleController.GetAllPermissionsByRoleExcludingDenied,
		)
	}
}
