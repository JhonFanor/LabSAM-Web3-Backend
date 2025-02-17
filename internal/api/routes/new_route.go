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
	news := nr.Router.Group("/news")
	{
		news.POST("/create", nr.ValidatorMiddleware.ValidateInput(&requests.NewsRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.CreateNews)
		news.GET("/get/all", nr.NewsController.GetAllNews)
		news.GET("/get/:id", nr.NewsController.GetNewsByID)
		news.PUT("/update/:id", nr.ValidatorMiddleware.ValidateInput(&requests.NewsUpdateRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.UpdateNews)
		news.DELETE("/delete/:id", nr.ValidatorMiddleware.ValidateInput(&requests.NewsUpdateRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.DeleteNews)
	}
}
