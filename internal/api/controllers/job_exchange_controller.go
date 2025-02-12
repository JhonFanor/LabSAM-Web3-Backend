package controllers

import (
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"

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
