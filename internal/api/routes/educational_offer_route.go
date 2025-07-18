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
		educationalOffer.GET("/user/me", eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.GetAllEducationalOffersByUserID)
		educationalOffer.GET("/admin/not-approved", eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.GetAllEducationalOffersNotApproved)
		educationalOffer.GET("/admin/not-approved/count", eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.CountEducationalOffersNotApproved)
		educationalOffer.GET("/:id", eor.TokenMiddleware.ValidateOptionalToken(), eor.EducationalOfferController.GetEducationalOfferByID)
		educationalOffer.PUT("/:id", eor.ValidatorMiddleware.ValidateInput(&requests.EducationalUpdateOfferRequest{}), eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.UpdateEducationalOffer)
		educationalOffer.PUT("/:id/approval", eor.ValidatorMiddleware.ValidateInput(&requests.ApprovalRequest{}), eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.SetEducationalOfferApproval)
		educationalOffer.DELETE("/:id", eor.TokenMiddleware.ValidateToken(), eor.EducationalOfferController.DeleteEducationalOffer)
	}
}
