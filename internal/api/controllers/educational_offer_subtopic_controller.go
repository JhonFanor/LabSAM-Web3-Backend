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

type EducationalOfferSubtopicControllerParams struct {
	fx.In
	EducationalOfferSubtopicService services.EducationalOfferSubtopicService
	EducationalOfferService         services.EducationalOfferService
}

type EducationalOfferSubtopicController struct {
	service                 services.EducationalOfferSubtopicService
	educationalOfferService services.EducationalOfferService
}

func NewEducationalOfferSubtopicController(p EducationalOfferSubtopicControllerParams) *EducationalOfferSubtopicController {
	return &EducationalOfferSubtopicController{
		service:                 p.EducationalOfferSubtopicService,
		educationalOfferService: p.EducationalOfferService,
	}
}

func (e *EducationalOfferSubtopicController) CreateEducationalOfferSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	educationalOfferSubtopicRequest := validatedInput.(*requests.EducationalOfferSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	educationalOffer, err := e.educationalOfferService.GetEducationalOfferByID(uint(id), claims.UserID, claims.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Educational Offer not found"})
		return
	}

	if educationalOffer.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range educationalOfferSubtopicRequest.SubtopicIDs {
		educationalOfferSubtopic := &models.EducationalOfferSubtopic{
			EducationalOfferID: uint(id),
			SubtopicID:         uint(subtopicID),
		}

		_, err := e.service.CreateEducationalOfferSubtopic(educationalOfferSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (e *EducationalOfferSubtopicController) DeleteEducationalOfferSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	educationalOfferSubtopicRequest := validatedInput.(*requests.EducationalOfferSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	educationalOffer, err := e.educationalOfferService.GetEducationalOfferByID(uint(id), claims.UserID, claims.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Educational Offer not found"})
		return
	}

	if educationalOffer.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range educationalOfferSubtopicRequest.SubtopicIDs {
		if err := e.service.DeleteEducationalOfferSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
