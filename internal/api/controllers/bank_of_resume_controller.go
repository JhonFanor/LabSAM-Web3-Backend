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
// @Description Create a new BankOfResume record. Requires authentication.
// @Tags BankOfResume
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.BankOfResumeCreateRequest true "Bank Of Resume Information"
// @Success 201 {object} models.BankOfResume "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /bank-of-resume/create [post]
func (b *BankOfResumeController) CreateBankOfResume(c *gin.Context) {
	validatedInput, _ := c.Get("input")
	bankOfResumeRequest := validatedInput.(*requests.BankOfResumeCreateRequest)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	var bankOfResume models.BankOfResume
	if err := mapstructure.Decode(bankOfResumeRequest, &bankOfResume); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
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
// @Description Fetch all BankOfResume records with pagination.
// @Tags BankOfResume
// @Produce json
// @Success 200 {object} dto.PaginationDTO "List of BankOfResume entries"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /bank-of-resume/get/all [get]
func (b *BankOfResumeController) GetAllBankOfResumes(c *gin.Context) {
	pagination, err := b.service.GetAllBankOfResumes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

// GetBankOfResumeByID godoc
// @Summary Get a single BankOfResume entry by ID
// @Description Fetch a single BankOfResume record by ID.
// @Tags BankOfResume
// @Produce json
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
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Bank Of Resume not found"})
		return
	}

	c.JSON(http.StatusOK, bankOfResume)
}

// UpdateBankOfResume godoc
// @Summary Update an existing BankOfResume entry
// @Description Update a BankOfResume record.
// @Tags BankOfResume
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path int true "Bank Of Resume ID"
// @Param input body requests.BankOfResumeUpdateRequest true "Updated Bank Of Resume Information"
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

	validatedInput, _ := c.Get("input")
	bankOfResumeRequest := validatedInput.(*requests.BankOfResumeUpdateRequest)
	var bankOfResume models.BankOfResume
	if err := mapstructure.Decode(bankOfResumeRequest, &bankOfResume); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	bankOfResume.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := b.service.UpdateBankOfResume(&bankOfResume, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bankOfResume)
}

// DeleteBankOfResume godoc
// @Summary Delete a BankOfResume entry
// @Description Delete a BankOfResume record by ID.
// @Tags BankOfResume
// @Produce json
// @Param Authorization header string true "Bearer Token"
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

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := b.service.DeleteBankOfResume(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
