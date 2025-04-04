package routes

import (
	"lamsam-web3-backend/internal/api/controllers"
	"lamsam-web3-backend/internal/api/middlewares"
	"lamsam-web3-backend/internal/dto/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type JobBoardRoutesParams struct {
	fx.In
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	JobBoardController  *controllers.JobBoardController
}

type JobBoardRoutes struct {
	Router              *gin.Engine
	ValidatorMiddleware *middlewares.ValidatorMiddleware
	TokenMiddleware     *middlewares.TokenMiddleware
	JobBoardController  *controllers.JobBoardController
}

func NewJobBoardRoutes(p JobBoardRoutesParams) *JobBoardRoutes {
	return &JobBoardRoutes{
		Router:              p.Router,
		ValidatorMiddleware: p.ValidatorMiddleware,
		TokenMiddleware:     p.TokenMiddleware,
		JobBoardController:  p.JobBoardController,
	}
}

func (jer *JobBoardRoutes) Routes() {
	jobBoard := jer.Router.Group("/job-exchange")
	{
		jobBoard.POST("/create", jer.ValidatorMiddleware.ValidateInput(&requests.JobBoardRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.CreateJobBoard)
		jobBoard.GET("/get/all", jer.JobBoardController.GetAllJobsExchange)
		jobBoard.GET("/get/:id", jer.JobBoardController.GetJobBoardByID)
		jobBoard.PUT("/update/:id", jer.ValidatorMiddleware.ValidateInput(&requests.JobBoardUpdateRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.UpdateJobBoard)
		jobBoard.DELETE("/delete/:id", jer.ValidatorMiddleware.ValidateInput(&requests.JobBoardUpdateRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.DeleteJobBoard)
	}
}
