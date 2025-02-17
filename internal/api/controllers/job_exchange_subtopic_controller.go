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

type JobExchangeSubtopicControllerParams struct {
	fx.In
	JobExchangeSubtopicService services.JobExchangeSubtopicService
	JobExchangeService         services.JobExchangeService
}

type JobExchangeSubtopicController struct {
	service            services.JobExchangeSubtopicService
	jobExchangeService services.JobExchangeService
}

func NewJobExchangeSubtopicController(p JobExchangeSubtopicControllerParams) *JobExchangeSubtopicController {
	return &JobExchangeSubtopicController{
		service:            p.JobExchangeSubtopicService,
		jobExchangeService: p.JobExchangeService,
	}
}

func (j *JobExchangeSubtopicController) CreateJobExchangeSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	jobExchangeSubtopicRequest := validatedInput.(*requests.JobExchangeSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	jobExchange, err := j.jobExchangeService.GetJobExchangeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Job Exchange not found"})
		return
	}

	if jobExchange.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range jobExchangeSubtopicRequest.SubtopicIDs {
		jobExchangeSubtopic := &models.JobExchangeSubtopic{
			JobExchangeID: uint(id),
			SubtopicID:    uint(subtopicID),
		}

		_, err := j.service.CreateJobExchangeSubtopic(jobExchangeSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (j *JobExchangeSubtopicController) DeleteJobExchangeSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	jobExchangeSubtopicRequest := validatedInput.(*requests.JobExchangeSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	jobExchange, err := j.jobExchangeService.GetJobExchangeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Job Exchange not found"})
		return
	}

	if jobExchange.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range jobExchangeSubtopicRequest.SubtopicIDs {
		if err := j.service.DeleteJobExchangeSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
