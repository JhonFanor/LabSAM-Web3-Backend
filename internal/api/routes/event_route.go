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
	Router               *gin.Engine
	ValidatorMiddleware  *middlewares.ValidatorMiddleware
	TokenMiddleware      *middlewares.TokenMiddleware
	PermissionMiddleware *middlewares.PermissionMiddleware
	EventController      *controllers.EventController
}

type EventRoutes struct {
	Router               *gin.Engine
	ValidatorMiddleware  *middlewares.ValidatorMiddleware
	TokenMiddleware      *middlewares.TokenMiddleware
	PermissionMiddleware *middlewares.PermissionMiddleware
	EventController      *controllers.EventController
}

func NewEventRoutes(p EventRoutesParams) *EventRoutes {
	return &EventRoutes{
		Router:               p.Router,
		ValidatorMiddleware:  p.ValidatorMiddleware,
		TokenMiddleware:      p.TokenMiddleware,
		PermissionMiddleware: p.PermissionMiddleware,
		EventController:      p.EventController,
	}
}

func (er *EventRoutes) Routes() {
	event := er.Router.Group("/api/event")
	{
		event.POST("", er.ValidatorMiddleware.ValidateInput(&requests.EventRequest{}), er.TokenMiddleware.ValidateToken(), er.PermissionMiddleware.RequirePermission("event:create"), er.EventController.CreateEvent)
		event.GET("", er.EventController.GetAllEvents)
		event.GET("/user/me", er.TokenMiddleware.ValidateToken(), er.EventController.GetAllEventsByUserID)
		event.GET("/admin/not-approved", er.TokenMiddleware.ValidateToken(), er.EventController.GetAllEventsNotApproved)
		event.GET("/admin/not-approved/count", er.TokenMiddleware.ValidateToken(), er.EventController.CountEventsNotApproved)
		event.GET("/count-by-subtopic", er.EventController.CountEventBySubtopic)
		event.GET("/:id", er.TokenMiddleware.ValidateOptionalToken(), er.EventController.GetEventByID)
		event.PUT("/:id", er.ValidatorMiddleware.ValidateInput(&requests.EventUpdateRequest{}), er.TokenMiddleware.ValidateToken(), er.EventController.UpdateEvent)
		event.PUT("/:id/approval", er.ValidatorMiddleware.ValidateInput(&requests.ApprovalRequest{}), er.TokenMiddleware.ValidateToken(), er.EventController.SetEventApproval)
		event.DELETE("/:id", er.TokenMiddleware.ValidateToken(), er.EventController.DeleteEvent)
	}
}
