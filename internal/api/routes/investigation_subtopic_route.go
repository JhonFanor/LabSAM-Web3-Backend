package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type InvestigationSubtopicRoutesParams struct {
	fx.In
	Router                          *gin.Engine
	ValidatorMiddleware             *middlewares.ValidatorMiddleware
	TonkenMiddleware                *middlewares.TokenMiddleware
	InvestigationSubtopicController *controllers.InvestigationSubtopicController
}

type InvestigationSubtopicRoutes struct {
	Router                          *gin.Engine
	ValidatorMiddleware             *middlewares.ValidatorMiddleware
	TonkenMiddleware                *middlewares.TokenMiddleware
	InvestigationSubtopicController *controllers.InvestigationSubtopicController
}

func NewInvestigationSubtopicRoutes(p InvestigationSubtopicRoutesParams) *InvestigationSubtopicRoutes {
	return &InvestigationSubtopicRoutes{
		Router:                          p.Router,
		ValidatorMiddleware:             p.ValidatorMiddleware,
		TonkenMiddleware:                p.TonkenMiddleware,
		InvestigationSubtopicController: p.InvestigationSubtopicController,
	}
}

func (ir *InvestigationSubtopicRoutes) Routes() {
	investigationSubtopic := ir.Router.Group("/api/investigation-subtopic")
	{
		investigationSubtopic.POST("/:id", ir.ValidatorMiddleware.ValidateInput(&requests.InvestigationSubtopicRequest{}), ir.TonkenMiddleware.ValidateToken(), ir.InvestigationSubtopicController.CreateInvestigationSubtopic)
		investigationSubtopic.DELETE("/:id", ir.ValidatorMiddleware.ValidateInput(&requests.InvestigationSubtopicRequest{}), ir.TonkenMiddleware.ValidateToken(), ir.InvestigationSubtopicController.DeleteInvestigationSubtopic)
	}
}
