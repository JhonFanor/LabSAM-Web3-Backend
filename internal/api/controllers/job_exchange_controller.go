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

type JobExchangeControllerParams struct {
	fx.In
	JobExchangeService services.JobExchangeService
}

type JobExchangeController struct {
	service services.JobExchangeService
}

func NewJobExchangeController(p JobExchangeControllerParams) *JobExchangeController {
	return &JobExchangeController{
		service: p.JobExchangeService,
	}
}

// CreateJobExchange godoc
// @Summary Create a new Job Exchange entry
// @Description This endpoint creates a new Job Exchange record. Requires authentication.
// @Tags JobExchange
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.JobExchangeRequest true "Job Exchange Information"
// @Success 201 {object} models.JobExchange "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /job-exchange/create [post]
func (j *JobExchangeController) CreateJobExchange(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	jobExchangeRequest := validatedInput.(*requests.JobExchangeRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var jobExchange models.JobExchange
	if err := mapstructure.Decode(jobExchangeRequest, &jobExchange); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdJobExchange, err := j.service.CreateJobExchange(&jobExchange, claims.UserID, jobExchangeRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdJobExchange)
}

func (j *JobExchangeController) GetAllJobsExchange(c *gin.Context) {
	pagination, err := j.service.GetAllJobsExchange(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

func (j *JobExchangeController) GetJobExchangeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	jobExchange, err := j.service.GetJobExchangeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Job Exchange not found"})
		return
	}

	c.JSON(http.StatusOK, jobExchange)
}

func (j *JobExchangeController) UpdateJobExchange(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	jobExchangeRequest := validatedInput.(*requests.JobExchangeUpdateRequest)
	var jobExchange models.JobExchange
	if err := mapstructure.Decode(jobExchangeRequest, &jobExchange); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	jobExchange.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := j.service.UpdateJobExchange(&jobExchange, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobExchange)
}

func (j *JobExchangeController) DeleteJobExchange(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := j.service.DeleteJobExchange(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
