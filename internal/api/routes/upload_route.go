package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UploadRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	UploadController    *controllers.UploadController
}

type UploadRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	UploadController    *controllers.UploadController
}

func NewUploadRoutes(p UploadRoutesParams) *UploadRoutes {
	return &UploadRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		UploadController:    p.UploadController,
	}
}

func (ur *UploadRoutes) Routes() {
	upload := ur.Router.Group("/api/upload")
	{
		upload.POST("/file", ur.UploadController.UploadFile)
		upload.DELETE("/:id", ur.UploadController.DeleteFile)
	}
}
