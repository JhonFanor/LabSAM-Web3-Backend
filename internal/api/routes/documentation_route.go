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
	documentation := dr.Router.Group("/api/documentation")
	{
		documentation.POST("", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.CreateDocumentation)
		documentation.GET("", dr.DocumentationController.GetAllDocumentations)
		documentation.GET("/user/me", dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.GetAllDocumentationsByUserID)
		documentation.GET("/admin/not-approved", dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.GetAllDocumentationsNotApproved)
		documentation.GET("/admin/not-approved/count", dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.CountDocumentationsNotApproved)
		documentation.GET("/count-by-subtopic", dr.DocumentationController.CountDocumentationBySubtopic)
		documentation.GET("/:id", dr.TokenMiddleware.ValidateOptionalToken(), dr.DocumentationController.GetDocumentationByID)
		documentation.PUT("/:id", dr.ValidatorMiddleware.ValidateInput(&requests.DocumentationUpdateRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.UpdateDocumentation)
		documentation.PUT("/:id/approval", dr.ValidatorMiddleware.ValidateInput(&requests.ApprovalRequest{}), dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.SetDocumentationApproval)
		documentation.DELETE("/:id", dr.TokenMiddleware.ValidateToken(), dr.DocumentationController.DeleteDocumentation)
	}
}
