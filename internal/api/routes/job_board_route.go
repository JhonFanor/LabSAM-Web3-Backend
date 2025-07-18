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
	jobBoard := jer.Router.Group("/api/job-board")
	{
		jobBoard.POST("", jer.ValidatorMiddleware.ValidateInput(&requests.JobBoardRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.CreateJobBoard)
		jobBoard.GET("", jer.JobBoardController.GetAllJobsBoard)
		jobBoard.GET("/user/me", jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.GetAllJobsBoardByUserID)
		jobBoard.GET("/admin/not-approved", jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.GetAllJobsBoardNotApproved)
		jobBoard.GET("/admin/not-approved/count", jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.CountJobsBoardNotApproved)
		jobBoard.GET("/:id", jer.TokenMiddleware.ValidateOptionalToken(), jer.JobBoardController.GetJobBoardByID)
		jobBoard.PUT("/:id", jer.ValidatorMiddleware.ValidateInput(&requests.JobBoardUpdateRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.UpdateJobBoard)
		jobBoard.PUT("/:id/approval", jer.ValidatorMiddleware.ValidateInput(&requests.ApprovalRequest{}), jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.SetJobBoardApproval)
		jobBoard.DELETE("/:id", jer.TokenMiddleware.ValidateToken(), jer.JobBoardController.DeleteJobBoard)
	}
}
