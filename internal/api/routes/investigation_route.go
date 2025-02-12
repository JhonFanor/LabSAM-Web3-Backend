package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type InvestigationRoutesParams struct {
	fx.In
	Router                  *gin.Engine
	ValidatorMiddleware     *middlewares.ValidatorMiddleware
	TokenMiddleware         *middlewares.TokenMiddleware
	InvestigationController *controllers.InvestigationController
}

type InvestigationRoutes struct {
	Router                  *gin.Engine
	ValidatorMiddleware     *middlewares.ValidatorMiddleware
	TokenMiddleware         *middlewares.TokenMiddleware
	InvestigationController *controllers.InvestigationController
}

func NewInvestigationRoutes(p InvestigationRoutesParams) *InvestigationRoutes {
	return &InvestigationRoutes{
		Router:                  p.Router,
		ValidatorMiddleware:     p.ValidatorMiddleware,
		TokenMiddleware:         p.TokenMiddleware,
		InvestigationController: p.InvestigationController,
	}
}

func (ir *InvestigationRoutes) Routes() {
	auth := ir.Router.Group("/investigation")
	{
		auth.POST("/create", ir.ValidatorMiddleware.ValidateInput(&requests.InvestigationRequest{}), ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.CreateInvestigation)
	}
}
