package controllers

import (
	"encoding/json"
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

func (b *BankOfResumeController) GetAllBankOfResumes(c *gin.Context) {
	pagination, err := b.service.GetAllBankOfResumes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.BankOfResumeGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (b *BankOfResumeController) GetBankOfResumeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	var userID uint
	var role string

	if claimsValue, exists := c.Get("claims"); exists {
		if claims, ok := claimsValue.(*security.Claims); ok {
			userID = claims.UserID
			role = claims.Role
		}
	}

	bankOfResume, err := b.service.GetBankOfResumeByID(uint(id), userID, role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Bank Of Resume not found"})
		return
	}

	var bankOfResumeResponse responses.BankOfResumeGetResponse
	if err := mapstructure.Decode(bankOfResume, &bankOfResumeResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, bankOfResumeResponse)
}

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
