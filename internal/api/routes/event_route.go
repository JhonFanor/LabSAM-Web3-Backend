package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type EventRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	EventController     *controllers.EventController
}

type EventRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	EventController     *controllers.EventController
}

func NewEventRoutes(p EventRoutesParams) *EventRoutes {
	return &EventRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		EventController:     p.EventController,
	}
}

func (er *EventRoutes) Routes() {
	auth := er.Router.Group("/event")
	{
		auth.POST("/create", er.ValidatorMiddleware.ValidateInput(&requests.EventRequest{}), er.TokenMiddleware.ValidateToken(), er.EventController.CreateEvent)
	}
}
