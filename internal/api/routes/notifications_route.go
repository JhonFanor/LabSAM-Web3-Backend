package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type NotificationRoutesParams struct {
	fx.In
	Router                 *gin.Engine
	TokenMiddleware        *middlewares.TokenMiddleware
	NotificationController *controllers.NotificationController
}

type NotificationRoutes struct {
	Router                 *gin.Engine
	TokenMiddleware        *middlewares.TokenMiddleware
	NotificationController *controllers.NotificationController
}

func NewNotificationRoutes(p NotificationRoutesParams) *NotificationRoutes {
	return &NotificationRoutes{
		Router:                 p.Router,
		TokenMiddleware:        p.TokenMiddleware,
		NotificationController: p.NotificationController,
	}
}

func (nr *NotificationRoutes) Routes() {
	notifications := nr.Router.Group("/api/notifications")
	{
		notifications.GET("", nr.TokenMiddleware.ValidateToken(), nr.NotificationController.GetAll)
		notifications.GET("/not-read/count", nr.TokenMiddleware.ValidateToken(), nr.NotificationController.CountNotificationNotRead)
		notifications.GET("/:id", nr.TokenMiddleware.ValidateToken(), nr.NotificationController.GetByID)
		notifications.PATCH("/:id/read", nr.TokenMiddleware.ValidateToken(), nr.NotificationController.MarkAsRead)
		notifications.PATCH("/read-all", nr.TokenMiddleware.ValidateToken(), nr.NotificationController.MarkAllAsRead)
		notifications.DELETE("/:id", nr.TokenMiddleware.ValidateToken(), nr.NotificationController.Delete)
	}
}
