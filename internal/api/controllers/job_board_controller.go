package controllers

import (
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type JobBoardControllerParams struct {
	fx.In
	JobBoardService services.JobBoardService
}

type JobBoardController struct {
	service services.JobBoardService
}

func NewJobBoardController(p JobBoardControllerParams) *JobBoardController {
	return &JobBoardController{
		service: p.JobBoardService,
	}
}

func (j *JobBoardController) CreateJobBoard(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	jobBoardRequest := validatedInput.(*requests.JobBoardRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var jobBoard models.JobBoard
	if err := mapstructure.Decode(jobBoardRequest, &jobBoard); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdJobBoard, err := j.service.CreateJobBoard(&jobBoard, claims.UserID, jobBoardRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdJobBoard)
}

func (j *JobBoardController) GetAllJobsExchange(c *gin.Context) {
	pagination, err := j.service.GetAllJobsExchange(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

func (j *JobBoardController) GetJobBoardByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	jobBoard, err := j.service.GetJobBoardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Job Exchange not found"})
		return
	}

	c.JSON(http.StatusOK, jobBoard)
}

func (j *JobBoardController) UpdateJobBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	jobBoardRequest := validatedInput.(*requests.JobBoardUpdateRequest)
	var jobBoard models.JobBoard
	if err := mapstructure.Decode(jobBoardRequest, &jobBoard); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	jobBoard.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := j.service.UpdateJobBoard(&jobBoard, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobBoard)
}

func (j *JobBoardController) DeleteJobBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := j.service.DeleteJobBoard(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
