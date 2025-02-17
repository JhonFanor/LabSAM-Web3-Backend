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

// BankOfResumeSubtopicControllerParams encapsulates dependencies for the controller
type BankOfResumeSubtopicControllerParams struct {
	fx.In
	BankOfResumeSubtopicService services.BankOfResumeSubtopicService
	BankOfResumeService         services.BankOfResumeService
}

// BankOfResumeSubtopicController handles requests related to BankOfResumeSubtopics
type BankOfResumeSubtopicController struct {
	service             services.BankOfResumeSubtopicService
	bankOfResumeService services.BankOfResumeService
}

// NewBankOfResumeSubtopicController initializes a new controller
func NewBankOfResumeSubtopicController(p BankOfResumeSubtopicControllerParams) *BankOfResumeSubtopicController {
	return &BankOfResumeSubtopicController{
		service:             p.BankOfResumeSubtopicService,
		bankOfResumeService: p.BankOfResumeService,
	}
}

// CreateBankOfResumeSubtopic assigns subtopics to a bank of resume
// @Summary Assign subtopics to a bank of resume
// @Description Associates a list of subtopics with a specific bank of resume
// @Tags BankOfResumeSubtopics
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "BankOfResume ID"
// @Param body body requests.BankOfResumeSubtopicRequest true "Subtopics to assign"
// @Success 201 {object} responses.SuccessResponse "Successfully assigned subtopics"
// @Failure 400 {object} responses.ErrorResponse "Invalid ID"
// @Failure 404 {object} responses.ErrorResponse "BankOfResume not found"
// @Failure 500 {object} responses.ErrorResponse "Internal Server Error"
// @Router /bank-of-resume/{id}/subtopics [post]
func (b *BankOfResumeSubtopicController) CreateBankOfResumeSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	bankOfResumeSubtopicRequest := validatedInput.(*requests.BankOfResumeSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	bankOfResume, err := b.bankOfResumeService.GetBankOfResumeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Bank Of Resume not found"})
		return
	}

	if bankOfResume.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range bankOfResumeSubtopicRequest.SubtopicIDs {
		bankOfResumeSubtopic := &models.BankOfResumeSubtopic{
			BankOfResumeID: uint(id),
			SubtopicID:     uint(subtopicID),
		}

		_, err := b.service.CreateBankOfResumeSubtopic(bankOfResumeSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

// DeleteBankOfResumeSubtopic removes subtopics from a bank of resume
// @Summary Remove subtopics from a bank of resume
// @Description Removes the association of subtopics from a specific bank of resume
// @Tags BankOfResumeSubtopics
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "BankOfResume ID"
// @Param body body requests.BankOfResumeSubtopicRequest true "Subtopics to remove"
// @Success 200 {object} responses.SuccessResponse "Successfully deleted"
// @Failure 400 {object} responses.ErrorResponse "Invalid ID"
// @Failure 404 {object} responses.ErrorResponse "BankOfResume not found"
// @Failure 500 {object} responses.ErrorResponse "Internal Server Error"
// @Router /bank-of-resume/{id}/subtopics [delete]
func (b *BankOfResumeSubtopicController) DeleteBankOfResumeSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	bankOfResumeSubtopicRequest := validatedInput.(*requests.BankOfResumeSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	bankOfResume, err := b.bankOfResumeService.GetBankOfResumeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Bank Of Resume not found"})
		return
	}

	if bankOfResume.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range bankOfResumeSubtopicRequest.SubtopicIDs {
		if err := b.service.DeleteBankOfResumeSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
