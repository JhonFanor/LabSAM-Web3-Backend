package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type NewsSubtopicRoutesParams struct {
	fx.In
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	TonkenMiddleware       *middlewares.TokenMiddleware
	NewsSubtopicController *controllers.NewsSubtopicController
}

type NewsSubtopicRoutes struct {
	Router                 *gin.Engine
	ValidatorMiddleware    *middlewares.ValidatorMiddleware
	TonkenMiddleware       *middlewares.TokenMiddleware
	NewsSubtopicController *controllers.NewsSubtopicController
}

func NewNewsSubtopicRoutes(p NewsSubtopicRoutesParams) *NewsSubtopicRoutes {
	return &NewsSubtopicRoutes{
		Router:                 p.Router,
		ValidatorMiddleware:    p.ValidatorMiddleware,
		TonkenMiddleware:       p.TonkenMiddleware,
		NewsSubtopicController: p.NewsSubtopicController,
	}
}

func (nr *NewsSubtopicRoutes) Routes() {
	newsSubtopic := nr.Router.Group("/news-subtopic")
	{
		newsSubtopic.POST("/create", nr.ValidatorMiddleware.ValidateInput(&requests.NewsSubtopicRequest{}), nr.TonkenMiddleware.ValidateToken(), nr.NewsSubtopicController.CreateNewsSubtopic)
		newsSubtopic.DELETE("/delete/:id", nr.ValidatorMiddleware.ValidateInput(&requests.NewsSubtopicRequest{}), nr.TonkenMiddleware.ValidateToken(), nr.NewsSubtopicController.DeleteNewsSubtopic)
	}
}
