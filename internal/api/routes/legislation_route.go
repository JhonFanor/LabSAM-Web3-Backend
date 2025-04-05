package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type LegislationRoutesParams struct {
	fx.In
	Router                *gin.Engine
	ValidatorMiddleware   *middlewares.ValidatorMiddleware
	TokenMiddleware       *middlewares.TokenMiddleware
	LegislationController *controllers.LegislationController
}

type LegislationRoutes struct {
	Router                *gin.Engine
	ValidatorMiddleware   *middlewares.ValidatorMiddleware
	TokenMiddleware       *middlewares.TokenMiddleware
	LegislationController *controllers.LegislationController
}

func NewLegislationRoutes(p LegislationRoutesParams) *LegislationRoutes {
	return &LegislationRoutes{
		Router:                p.Router,
		ValidatorMiddleware:   p.ValidatorMiddleware,
		TokenMiddleware:       p.TokenMiddleware,
		LegislationController: p.LegislationController,
	}
}

func (lr *LegislationRoutes) Routes() {
	legislation := lr.Router.Group("/api/legislation")
	{
		legislation.POST("", lr.ValidatorMiddleware.ValidateInput(&requests.LegislationRequest{}), lr.TokenMiddleware.ValidateToken(), lr.LegislationController.CreateLegislation)
		legislation.GET("", lr.LegislationController.GetAllLegislations)
		legislation.GET("/:id", lr.LegislationController.GetLegislationByID)
		legislation.PUT("/:id", lr.ValidatorMiddleware.ValidateInput(&requests.LegislationUpdateRequest{}), lr.TokenMiddleware.ValidateToken(), lr.LegislationController.UpdateLegislation)
		legislation.PUT("/:id", lr.ValidatorMiddleware.ValidateInput(&requests.LegislationUpdateRequest{}), lr.TokenMiddleware.ValidateToken(), lr.LegislationController.DeleteLegislation)
	}
}
