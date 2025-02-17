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

type LegislationSubtopicControllerParams struct {
	fx.In
	LegislationSubtopicService services.LegislationSubtopicService
	LegislationService         services.LegislationService
}

type LegislationSubtopicController struct {
	service            services.LegislationSubtopicService
	legislationService services.LegislationService
}

func NewLegislationSubtopicController(p LegislationSubtopicControllerParams) *LegislationSubtopicController {
	return &LegislationSubtopicController{
		service:            p.LegislationSubtopicService,
		legislationService: p.LegislationService,
	}
}

func (l *LegislationSubtopicController) CreateLegislationSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	legislationSubtopicRequest := validatedInput.(*requests.LegislationSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	legislation, err := l.legislationService.GetLegislationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Legislation not found"})
		return
	}

	if legislation.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range legislationSubtopicRequest.SubtopicIDs {
		legislationSubtopic := &models.LegislationSubtopic{
			LegislationID: uint(id),
			SubtopicID:    uint(subtopicID),
		}

		_, err := l.service.CreateLegislationSubtopic(legislationSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (l *LegislationSubtopicController) DeleteLegislationSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	legislationSubtopicRequest := validatedInput.(*requests.LegislationSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	legislation, err := l.legislationService.GetLegislationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Legislation not found"})
		return
	}

	if legislation.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range legislationSubtopicRequest.SubtopicIDs {
		if err := l.service.DeleteLegislationSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
