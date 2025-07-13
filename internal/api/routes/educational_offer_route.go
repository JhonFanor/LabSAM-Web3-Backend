package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type EducationalOfferRoutesParams struct {
	fx.In
	Router                     *gin.Engine
	ValidatorMiddleware        *middlewares.ValidatorMiddleware
	TokenMiddleware            *middlewares.TokenMiddleware
	EducationalOfferController *controllers.EducationalOfferController
}

type EducationalOfferRoutes struct {
	Router                     *gin.Engine
	ValidatorMiddleware        *middlewares.ValidatorMiddleware
	TokenMiddleware            *middlewares.TokenMiddleware
	EducationalOfferController *controllers.EducationalOfferController
}

func NewEducationalOfferRoutes(p EducationalOfferRoutesParams) *EducationalOfferRoutes {
	return &EducationalOfferRoutes{
		Router:                     p.Router,
		ValidatorMiddleware:        p.ValidatorMiddleware,
		TokenMiddleware:            p.TokenMiddleware,
		EducationalOfferController: p.EducationalOfferController,
	}
}

func (eor *EducationalOfferRoutes) Routes() {
	educationalOffer := eor.Router.Group("/api/educational-offer")
	{
		educationalOffer.POST("", eor.ValidatorMiddleware.ValidateInput(&requests.EducationalOfferRequest{}), eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.CreateEducationalOffer)
		educationalOffer.GET("", eor.EducationalOfferController.GetAllEducationalOffers)
		educationalOffer.GET("/:id", eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.GetEducationalOfferByID)
		educationalOffer.PUT("/:id", eor.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeUpdateRequest{}), eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.UpdateEducationalOffer)
		educationalOffer.DELETE("/:id", eor.ValidatorMiddleware.ValidateInput(&requests.BankOfResumeUpdateRequest{}), eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.DeleteEducationalOffer)
	}
}
