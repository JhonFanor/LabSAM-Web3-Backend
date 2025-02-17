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

type DocumentationSubtopicControllerParams struct {
	fx.In
	DocumentationSubtopicService services.DocumentationSubtopicService
	DocumentationService         services.DocumentationService
}

type DocumentationSubtopicController struct {
	service              services.DocumentationSubtopicService
	documentationService services.DocumentationService
}

func NewDocumentationSubtopicController(p DocumentationSubtopicControllerParams) *DocumentationSubtopicController {
	return &DocumentationSubtopicController{
		service:              p.DocumentationSubtopicService,
		documentationService: p.DocumentationService,
	}
}

func (d *DocumentationSubtopicController) CreateDocumentationSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	documentationSubtopicRequest := validatedInput.(*requests.DocumentationSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	documentation, err := d.documentationService.GetDocumentationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Documentation not found"})
		return
	}

	if documentation.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range documentationSubtopicRequest.SubtopicIDs {
		documentationSubtopic := &models.DocumentationSubtopic{
			DocumentationID: uint(id),
			SubtopicID:      uint(subtopicID),
		}

		_, err := d.service.CreateDocumentationSubtopic(documentationSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (d *DocumentationSubtopicController) DeleteDocumentationSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	documentationSubtopicRequest := validatedInput.(*requests.DocumentationSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	documentation, err := d.documentationService.GetDocumentationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Documentation not found"})
		return
	}

	if documentation.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range documentationSubtopicRequest.SubtopicIDs {
		if err := d.service.DeleteDocumentationSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
