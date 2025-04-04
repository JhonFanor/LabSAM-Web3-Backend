package controllers

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type JobBoardSubtopicControllerParams struct {
	fx.In
	JobBoardSubtopicService services.JobBoardSubtopicService
	JobBoardService         services.JobBoardService
}

type JobBoardSubtopicController struct {
	service         services.JobBoardSubtopicService
	jobBoardService services.JobBoardService
}

func NewJobBoardSubtopicController(p JobBoardSubtopicControllerParams) *JobBoardSubtopicController {
	return &JobBoardSubtopicController{
		service:         p.JobBoardSubtopicService,
		jobBoardService: p.JobBoardService,
	}
}

func (j *JobBoardSubtopicController) CreateJobBoardSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	jobBoardSubtopicRequest := validatedInput.(*requests.JobBoardSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	jobBoard, err := j.jobBoardService.GetJobBoardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Job Exchange not found"})
		return
	}

	if jobBoard.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range jobBoardSubtopicRequest.SubtopicIDs {
		jobBoardSubtopic := &models.JobBoardSubtopic{
			JobBoardID: uint(id),
			SubtopicID: uint(subtopicID),
		}

		_, err := j.service.CreateJobBoardSubtopic(jobBoardSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (j *JobBoardSubtopicController) DeleteJobBoardSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	jobBoardSubtopicRequest := validatedInput.(*requests.JobBoardSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	jobBoard, err := j.jobBoardService.GetJobBoardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Job Exchange not found"})
		return
	}

	if jobBoard.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range jobBoardSubtopicRequest.SubtopicIDs {
		if err := j.service.DeleteJobBoardSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
