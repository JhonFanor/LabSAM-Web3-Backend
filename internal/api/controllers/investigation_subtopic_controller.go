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

type InvestigationSubtopicControllerParams struct {
	fx.In
	InvestigationSubtopicService services.InvestigationSubtopicService
	InvestigationService         services.InvestigationService
}

type InvestigationSubtopicController struct {
	service              services.InvestigationSubtopicService
	investigationService services.InvestigationService
}

func NewInvestigationSubtopicController(p InvestigationSubtopicControllerParams) *InvestigationSubtopicController {
	return &InvestigationSubtopicController{
		service:              p.InvestigationSubtopicService,
		investigationService: p.InvestigationService,
	}
}

func (i *InvestigationSubtopicController) CreateInvestigationSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	investigationSubtopicRequest := validatedInput.(*requests.InvestigationSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	investigation, err := i.investigationService.GetInvestigationByID(uint(id), claims.UserID, claims.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Investigation not found"})
		return
	}

	if investigation.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range investigationSubtopicRequest.SubtopicIDs {
		investigationSubtopic := &models.InvestigationSubtopic{
			InvestigationID: uint(id),
			SubtopicID:      uint(subtopicID),
		}

		_, err := i.service.CreateInvestigationSubtopic(investigationSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (i *InvestigationSubtopicController) DeleteInvestigationSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	investigationSubtopicRequest := validatedInput.(*requests.InvestigationSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	investigation, err := i.investigationService.GetInvestigationByID(uint(id), claims.UserID, claims.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Investigation not found"})
		return
	}

	if investigation.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range investigationSubtopicRequest.SubtopicIDs {
		if err := i.service.DeleteInvestigationSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
