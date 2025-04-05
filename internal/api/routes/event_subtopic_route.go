package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type EventSubtopicRoutesParams struct {
	fx.In
	Router                  *gin.Engine
	ValidatorMiddleware     *middlewares.ValidatorMiddleware
	TonkenMiddleware        *middlewares.TokenMiddleware
	EventSubtopicController *controllers.EventSubtopicController
}

type EventSubtopicRoutes struct {
	Router                  *gin.Engine
	ValidatorMiddleware     *middlewares.ValidatorMiddleware
	TonkenMiddleware        *middlewares.TokenMiddleware
	EventSubtopicController *controllers.EventSubtopicController
}

func NewEventSubtopicRoutes(p EventSubtopicRoutesParams) *EventSubtopicRoutes {
	return &EventSubtopicRoutes{
		Router:                  p.Router,
		ValidatorMiddleware:     p.ValidatorMiddleware,
		TonkenMiddleware:        p.TonkenMiddleware,
		EventSubtopicController: p.EventSubtopicController,
	}
}

func (er *EventSubtopicRoutes) Routes() {
	eventSubtopic := er.Router.Group("/api/event-subtopic")
	{
		eventSubtopic.POST("/:id", er.ValidatorMiddleware.ValidateInput(&requests.EventSubtopicRequest{}), er.TonkenMiddleware.ValidateToken(), er.EventSubtopicController.CreateEventSubtopic)
		eventSubtopic.DELETE("/:id", er.ValidatorMiddleware.ValidateInput(&requests.EventSubtopicRequest{}), er.TonkenMiddleware.ValidateToken(), er.EventSubtopicController.DeleteEventSubtopic)
	}
}
