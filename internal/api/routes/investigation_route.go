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
	investigation := ir.Router.Group("/api/investigation")
	{
		investigation.POST("", ir.ValidatorMiddleware.ValidateInput(&requests.InvestigationRequest{}), ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.CreateInvestigation)
		investigation.GET("", ir.InvestigationController.GetAllInvestigations)
		investigation.GET("/user/me", ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.GetAllInvestigationsByUserID)
		investigation.GET("/admin/not-approved", ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.GetAllInvestigationsNotApproved)
		investigation.GET("/admin/not-approved/count", ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.CountInvestigationsNotApproved)
		investigation.GET("/:id", ir.TokenMiddleware.ValidateOptionalToken(), ir.InvestigationController.GetInvestigationByID)
		investigation.PUT("/:id", ir.ValidatorMiddleware.ValidateInput(&requests.InvestigationUpdateRequest{}), ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.UpdateInvestigation)
		investigation.DELETE("/:id", ir.ValidatorMiddleware.ValidateInput(&requests.InvestigationUpdateRequest{}), ir.TokenMiddleware.ValidateToken(), ir.InvestigationController.DeleteInvestigation)
	}
}
