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
	documentation := dr.Router.Group("/documentation")
	{
		documentation.POST("/create", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.CreateDocumentation)
		documentation.GET("/get/all", dr.DocumentationController.GetAllDocumentations)
		documentation.GET("/get/:id", dr.DocumentationController.GetDocumentationByID)
		documentation.PUT("/update/:id", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationUpdateRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.UpdateDocumentation)
		documentation.DELETE("/delete/:id", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationUpdateRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.DeleteDocumentation)
	}
}
