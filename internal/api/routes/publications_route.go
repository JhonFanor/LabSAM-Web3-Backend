package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PublicationRoutesParams struct {
	fx.In
	Router                *gin.Engine
	TokenMiddleware       *middlewares.TokenMiddleware
	PublicationController *controllers.PublicationsController
}

type PublicationRoutes struct {
	Router                *gin.Engine
	TokenMiddleware       *middlewares.TokenMiddleware
	PublicationController *controllers.PublicationsController
}

func NewPublicationRoutes(p PublicationRoutesParams) *PublicationRoutes {
	return &PublicationRoutes{
		Router:                p.Router,
		TokenMiddleware:       p.TokenMiddleware,
		PublicationController: p.PublicationController,
	}
}

func (pr *PublicationRoutes) Routes() {
	publications := pr.Router.Group("/api/publications")
	{
		publications.GET("", pr.TokenMiddleware.ValidateOptionalToken(), pr.PublicationController.GetAllPublications)
	}
}
