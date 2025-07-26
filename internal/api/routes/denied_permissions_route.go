package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type DeniedPermissionUserRoutesParams struct {
	fx.In
	Router                         *gin.Engine
	TokenMiddleware                *middlewares.TokenMiddleware
	DeniedPermissionUserController *controllers.DeniedPermissionUserController
}

type DeniedPermissionUserRoutes struct {
	Router                         *gin.Engine
	TokenMiddleware                *middlewares.TokenMiddleware
	DeniedPermissionUserController *controllers.DeniedPermissionUserController
}

func NewDeniedPermissionUserRoutes(p DeniedPermissionUserRoutesParams) *DeniedPermissionUserRoutes {
	return &DeniedPermissionUserRoutes{
		Router:                         p.Router,
		TokenMiddleware:                p.TokenMiddleware,
		DeniedPermissionUserController: p.DeniedPermissionUserController,
	}
}

func (r *DeniedPermissionUserRoutes) Routes() {
	denied := r.Router.Group("/api/denied-permissions")
	{
		denied.GET("/user/:user_id", r.TokenMiddleware.ValidateToken(), r.DeniedPermissionUserController.GetAllDeniedPermissionsByUser)

		denied.POST("/assign", r.TokenMiddleware.ValidateToken(), r.DeniedPermissionUserController.AssignDeniedPermission)

		denied.POST("/revoke", r.TokenMiddleware.ValidateToken(), r.DeniedPermissionUserController.RevokeDeniedPermission)
	}
}
