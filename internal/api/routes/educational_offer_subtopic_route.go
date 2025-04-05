package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type EducationalOfferSubtopicRoutesParams struct {
	fx.In
	Router                             *gin.Engine
	ValidatorMiddleware                *middlewares.ValidatorMiddleware
	TonkenMiddleware                   *middlewares.TokenMiddleware
	EducationalOfferSubtopicController *controllers.EducationalOfferSubtopicController
}

type EducationalOfferSubtopicRoutes struct {
	Router                             *gin.Engine
	ValidatorMiddleware                *middlewares.ValidatorMiddleware
	TonkenMiddleware                   *middlewares.TokenMiddleware
	EducationalOfferSubtopicController *controllers.EducationalOfferSubtopicController
}

func NewEducationalOfferSubtopicRoutes(p EducationalOfferSubtopicRoutesParams) *EducationalOfferSubtopicRoutes {
	return &EducationalOfferSubtopicRoutes{
		Router:                             p.Router,
		ValidatorMiddleware:                p.ValidatorMiddleware,
		TonkenMiddleware:                   p.TonkenMiddleware,
		EducationalOfferSubtopicController: p.EducationalOfferSubtopicController,
	}
}

func (er *EducationalOfferSubtopicRoutes) Routes() {
	educationalOfferSubtopic := er.Router.Group("/api/educational-offer-subtopic")
	{
		educationalOfferSubtopic.POST("/:id", er.ValidatorMiddleware.ValidateInput(&requests.EducationalOfferSubtopicRequest{}), er.TonkenMiddleware.ValidateToken(), er.EducationalOfferSubtopicController.CreateEducationalOfferSubtopic)
		educationalOfferSubtopic.DELETE("/:id", er.ValidatorMiddleware.ValidateInput(&requests.EducationalOfferSubtopicRequest{}), er.TonkenMiddleware.ValidateToken(), er.EducationalOfferSubtopicController.DeleteEducationalOfferSubtopic)
	}
}
