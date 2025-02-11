package controllers

import (
	"context"
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

type BankOfResumeControllerParams struct {
	fx.In
	BankOfResumeService services.BankOfResumeService
}

type BankOfResumeController struct {
	service services.BankOfResumeService
}

func NewBankOfResumeController(p BankOfResumeControllerParams) *BankOfResumeController {
	return &BankOfResumeController{
		service: p.BankOfResumeService,
	}
}

// CreateBankOfResume godoc
// @Summary Create a new BankOfResume entry
// @Description This endpoint creates a new BankOfResume record. Requires authentication.
// @Tags BankOfResume
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.BankOfResumeRequest true "Bank Of Resume Information"
// @Success 201 {object} models.BankOfResume "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /bank-of-resume/create [post]
func (b *BankOfResumeController) CreateBankOfResume(c *gin.Context) {

	validatedInput, _ := c.Get("input")

	bankOfResumeRequest := validatedInput.(*requests.BankOfResumeRequest)

	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var bankOfResume models.BankOfResume
	if err := mapstructure.Decode(bankOfResumeRequest, &bankOfResume); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdBankOfResume, err := b.service.CreateBankOfResume(&bankOfResume, claims.UserID, bankOfResumeRequest.SubtopicIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBankOfResume)
}

// GetAllBankOfResumes godoc
// @Summary Get all BankOfResume entries
// @Description This endpoint fetches all BankOfResume records with pagination.
// @Tags BankOfResume
// @Produce  json
// @Success 200 {object} []models.BankOfResume "List of BankOfResume entries"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /bank-of-resume/get-all [get]
func (b *BankOfResumeController) GetAllBankOfResumes(c *gin.Context) {
	ctx := context.Background()
	data, pagination, err := b.service.GetAllBankOfResumes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

// GetBankOfResumeByID godoc
// @Summary Get a single BankOfResume entry by ID
// @Description This endpoint fetches a single BankOfResume record.
// @Tags BankOfResume
// @Produce  json
// @Param id path int true "Bank Of Resume ID"
// @Success 200 {object} models.BankOfResume "Successfully retrieved"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 404 {object} responses.ErrorResponse "Not found"
// @Router /bank-of-resume/get/{id} [get]
func (b *BankOfResumeController) GetBankOfResumeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	bankOfResume, err := b.service.GetBankOfResumeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "BankOfResume not found"})
		return
	}

	c.JSON(http.StatusOK, bankOfResume)
}

// UpdateBankOfResume godoc
// @Summary Update an existing BankOfResume entry
// @Description This endpoint updates a BankOfResume record.
// @Tags BankOfResume
// @Accept  json
// @Produce  json
// @Param id path int true "Bank Of Resume ID"
// @Param input body models.BankOfResume true "Updated Bank Of Resume Information"
// @Success 200 {object} models.BankOfResume "Successfully updated"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 404 {object} responses.ErrorResponse "Not found"
// @Router /bank-of-resume/update/{id} [put]
func (b *BankOfResumeController) UpdateBankOfResume(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	var bankOfResume models.BankOfResume
	if err := c.ShouldBindJSON(&bankOfResume); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	bankOfResume.ID = uint(id)
	if err := b.service.UpdateBankOfResume(&bankOfResume); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bankOfResume)
}

// DeleteBankOfResume godoc
// @Summary Delete a BankOfResume entry
// @Description This endpoint deletes a BankOfResume record.
// @Tags BankOfResume
// @Produce  json
// @Param id path int true "Bank Of Resume ID"
// @Success 200 {object} responses.SuccessResponse "Successfully deleted"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 404 {object} responses.ErrorResponse "Not found"
// @Router /bank-of-resume/delete/{id} [delete]
func (b *BankOfResumeController) DeleteBankOfResume(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	if err := b.service.DeleteBankOfResume(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
