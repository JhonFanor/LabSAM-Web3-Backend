package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type DocumentationSubtopicRoutesParams struct {
	fx.In
	Router                          *gin.Engine
	ValidatorMiddleware             *middlewares.ValidatorMiddleware
	TonkenMiddleware                *middlewares.TokenMiddleware
	DocumentationSubtopicController *controllers.DocumentationSubtopicController
}

type DocumentationSubtopicRoutes struct {
	Router                          *gin.Engine
	ValidatorMiddleware             *middlewares.ValidatorMiddleware
	TonkenMiddleware                *middlewares.TokenMiddleware
	DocumentationSubtopicController *controllers.DocumentationSubtopicController
}

func NewDocumentationSubtopicRoutes(p DocumentationSubtopicRoutesParams) *DocumentationSubtopicRoutes {
	return &DocumentationSubtopicRoutes{
		Router:                          p.Router,
		ValidatorMiddleware:             p.ValidatorMiddleware,
		TonkenMiddleware:                p.TonkenMiddleware,
		DocumentationSubtopicController: p.DocumentationSubtopicController,
	}
}

func (dr *DocumentationSubtopicRoutes) Routes() {
	documentationSubtopic := dr.Router.Group("/api/documentation-subtopic")
	{
		documentationSubtopic.POST("/:id", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationSubtopicRequest{}), dr.TonkenMiddleware.ValidateToken(), dr.DocumentationSubtopicController.CreateDocumentationSubtopic)
		documentationSubtopic.DELETE("/:id", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationSubtopicRequest{}), dr.TonkenMiddleware.ValidateToken(), dr.DocumentationSubtopicController.DeleteDocumentationSubtopic)
	}
}
