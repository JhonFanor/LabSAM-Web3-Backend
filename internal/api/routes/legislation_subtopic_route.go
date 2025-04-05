package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type LegislationSubtopicRoutesParams struct {
	fx.In
	Router                        *gin.Engine
	ValidatorMiddleware           *middlewares.ValidatorMiddleware
	TonkenMiddleware              *middlewares.TokenMiddleware
	LegislationSubtopicController *controllers.LegislationSubtopicController
}

type LegislationSubtopicRoutes struct {
	Router                        *gin.Engine
	ValidatorMiddleware           *middlewares.ValidatorMiddleware
	TonkenMiddleware              *middlewares.TokenMiddleware
	LegislationSubtopicController *controllers.LegislationSubtopicController
}

func NewLegislationSubtopicRoutes(p LegislationSubtopicRoutesParams) *LegislationSubtopicRoutes {
	return &LegislationSubtopicRoutes{
		Router:                        p.Router,
		ValidatorMiddleware:           p.ValidatorMiddleware,
		TonkenMiddleware:              p.TonkenMiddleware,
		LegislationSubtopicController: p.LegislationSubtopicController,
	}
}

func (lr *LegislationSubtopicRoutes) Routes() {
	legislationSubtopic := lr.Router.Group("/api/legislation-subtopic")
	{
		legislationSubtopic.POST("/:id", lr.ValidatorMiddleware.ValidateInput(&requests.LegislationSubtopicRequest{}), lr.TonkenMiddleware.ValidateToken(), lr.LegislationSubtopicController.CreateLegislationSubtopic)
		legislationSubtopic.DELETE("/:id", lr.ValidatorMiddleware.ValidateInput(&requests.LegislationSubtopicRequest{}), lr.TonkenMiddleware.ValidateToken(), lr.LegislationSubtopicController.DeleteLegislationSubtopic)
	}
}
