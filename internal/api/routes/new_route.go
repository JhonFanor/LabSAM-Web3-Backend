package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type NewsRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	NewsController      *controllers.NewsController
}

type NewsRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	NewsController      *controllers.NewsController
}

func NewNewsRoutes(p NewsRoutesParams) *NewsRoutes {
	return &NewsRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		NewsController:      p.NewsController,
	}
}

func (nr *NewsRoutes) Routes() {
	news := nr.Router.Group("/aṕi/news")
	{
		news.POST("", nr.ValidatorMiddleware.ValidateInput(&requests.NewsRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.CreateNews)
		news.GET("", nr.NewsController.GetAllNews)
		news.GET("/:id", nr.NewsController.GetNewsByID)
		news.PUT("/:id", nr.ValidatorMiddleware.ValidateInput(&requests.NewsUpdateRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.UpdateNews)
		news.DELETE("/:id", nr.TokenMiddleware.ValidateToken(), nr.NewsController.DeleteNews)
	}
}
