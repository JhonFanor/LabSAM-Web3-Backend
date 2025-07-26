package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PermissionRoutesParams struct {
	fx.In
	Router               *gin.Engine
	TokenMiddleware      *middlewares.TokenMiddleware
	ValidatorMiddleware  *middlewares.ValidatorMiddleware
	PermissionController *controllers.PermissionController
}

type PermissionRoutes struct {
	Router               *gin.Engine
	TokenMiddleware      *middlewares.TokenMiddleware
	ValidatorMiddleware  *middlewares.ValidatorMiddleware
	PermissionController *controllers.PermissionController
}

func NewPermissionRoutes(p PermissionRoutesParams) *PermissionRoutes {
	return &PermissionRoutes{
		Router:               p.Router,
		TokenMiddleware:      p.TokenMiddleware,
		ValidatorMiddleware:  p.ValidatorMiddleware,
		PermissionController: p.PermissionController,
	}
}

func (r *PermissionRoutes) Routes() {
	permissions := r.Router.Group("/api/permissions")
	{
		permissions.POST("/assignable",
			r.ValidatorMiddleware.ValidateInput(&requests.PermissionRequest{}),
			r.TokenMiddleware.ValidateToken(),
			r.PermissionController.GetAllAssignablePermissionsToUser,
		)
	}
}
