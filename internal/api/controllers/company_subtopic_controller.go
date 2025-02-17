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

type CompanySubtopicControllerParams struct {
	fx.In
	CompanySubtopicService services.CompanySubtopicService
	CompanyService         services.CompanyService
}

type CompanySubtopicController struct {
	service        services.CompanySubtopicService
	companyService services.CompanyService
}

func NewComapanySubtopicController(p CompanySubtopicControllerParams) *CompanySubtopicController {
	return &CompanySubtopicController{
		service:        p.CompanySubtopicService,
		companyService: p.CompanyService,
	}
}

func (c *CompanySubtopicController) CreateCompanySubtopic(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := context.Get("input")
	companySubtopicRequest := validatedInput.(*requests.CompanySubtopicRequest)

	claimsValue, _ := context.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	company, err := c.companyService.GetCompanyByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Company not found"})
		return
	}

	if company.UserID != claims.UserID && claims.Role != "admin" {
		context.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range companySubtopicRequest.SubtopicIDs {
		companySubtopic := &models.CompanySubtopic{
			CompanyID:  uint(id),
			SubtopicID: uint(subtopicID),
		}

		_, err := c.service.CreateCompanySubtopic(companySubtopic)
		if err != nil {
			context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	context.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (c *CompanySubtopicController) DeleteCompanySubtopic(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := context.Get("input")
	companySubtopicRequest := validatedInput.(*requests.CompanySubtopicRequest)

	claimsValue, _ := context.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	company, err := c.companyService.GetCompanyByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Company not found"})
		return
	}

	if company.UserID != claims.UserID && claims.Role != "admin" {
		context.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range companySubtopicRequest.SubtopicIDs {
		if err := c.service.DeleteCompanySubtopic(uint(id), subtopicID); err != nil {
			context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	context.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
