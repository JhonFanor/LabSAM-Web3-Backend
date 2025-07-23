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
	news := nr.Router.Group("/api/news")
	{
		news.POST("", nr.ValidatorMiddleware.ValidateInput(&requests.NewsRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.CreateNews)
		news.GET("", nr.NewsController.GetAllNews)
		news.GET("/user/me", nr.TokenMiddleware.ValidateToken(), nr.NewsController.GetAllNewsByUserID)
		news.GET("/admin/not-approved", nr.TokenMiddleware.ValidateToken(), nr.NewsController.GetAllNewsNotApproved)
		news.GET("/admin/not-approved/count", nr.TokenMiddleware.ValidateToken(), nr.NewsController.CountNewsNotApproved)
		news.GET("/count-by-subtopic", nr.NewsController.CountNewsBySubtopic)
		news.GET("/:id", nr.TokenMiddleware.ValidateOptionalToken(), nr.NewsController.GetNewsByID)
		news.PUT("/:id", nr.ValidatorMiddleware.ValidateInput(&requests.NewsUpdateRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.UpdateNews)
		news.PUT("/:id/approval", nr.ValidatorMiddleware.ValidateInput(&requests.ApprovalRequest{}), nr.TokenMiddleware.ValidateToken(), nr.NewsController.SetNewsApproval)
		news.DELETE("/:id", nr.TokenMiddleware.ValidateToken(), nr.NewsController.DeleteNews)
	}
}
