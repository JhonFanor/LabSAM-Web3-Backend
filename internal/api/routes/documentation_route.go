package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type DocumentationRoutesParams struct {
	fx.In
	Router                  *gin.Engine
	ValidatorMiddleware     *middlewares.ValidatorMiddleware
	TokenMiddleware         *middlewares.TokenMiddleware
	DocumentationController *controllers.DocumentationController
}

type DocumentationRoutes struct {
	Router                  *gin.Engine
	ValidatorMiddleware     *middlewares.ValidatorMiddleware
	TokenMiddleware         *middlewares.TokenMiddleware
	DocumentationController *controllers.DocumentationController
}

func NewDocumentationRoutes(p DocumentationRoutesParams) *DocumentationRoutes {
	return &DocumentationRoutes{
		Router:                  p.Router,
		ValidatorMiddleware:     p.ValidatorMiddleware,
		TokenMiddleware:         p.TokenMiddleware,
		DocumentationController: p.DocumentationController,
	}
}

func (dr *DocumentationRoutes) Routes() {
	auth := dr.Router.Group("/documentation")
	{
		auth.POST("/create", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.CreateDocumentation)
	}
}
