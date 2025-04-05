package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type TopicRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	TopicController     *controllers.TopicController
}

type TopicRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	TopicController     *controllers.TopicController
}

func NewTopicRoutes(p TopicRoutesParams) *TopicRoutes {
	return &TopicRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		TopicController:     p.TopicController,
	}
}

func (tr *TopicRoutes) Routes() {
	topic := tr.Router.Group("/api/topic")
	{
		topic.POST("", tr.ValidatorMiddleware.ValidateInput(&requests.TopicRequest{}), tr.TokenMiddleware.ValidateToken(), tr.TopicController.CreateTopic)
		topic.GET("", tr.TopicController.GetAllTopics)
		topic.GET("/:id", tr.TopicController.GetTopicByID)
		topic.PUT("/:id", tr.ValidatorMiddleware.ValidateInput(&requests.TopicRequest{}), tr.TokenMiddleware.ValidateToken(), tr.TopicController.UpdateTopic)
		topic.DELETE("/:id", tr.TokenMiddleware.ValidateToken(), tr.TopicController.DeleteTopic)
	}
}
