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

type CompanyControllerParams struct {
	fx.In
	CompanyService      services.CompanyService
	LocalitationService services.LocalitationService
}

type CompanyController struct {
	service             services.CompanyService
	localitationService services.LocalitationService
}

func NewCompanyController(p CompanyControllerParams) *CompanyController {
	return &CompanyController{
		service:             p.CompanyService,
		localitationService: p.LocalitationService,
	}
}

func (c *CompanyController) CreateCompany(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	companyRequest := validatedInput.(*requests.CompanyRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var company models.Company
	if err := mapstructure.Decode(companyRequest, &company); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	company.LocalitationID = nil

	if companyRequest.Localiatation != nil {
		var localitationRequest models.Localitation
		if err := mapstructure.Decode(companyRequest.Localiatation, &localitationRequest); err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
			return
		}

		localitation, err := c.localitationService.AssignLocalitation(&localitationRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}

		company.LocalitationID = &localitation.ID
	}

	createdCompany, err := c.service.CreateCompany(&company, claims.UserID, companyRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdCompany)
}

func (c *CompanyController) GetAllCompanies(context *gin.Context) {
	pagination, err := c.service.GetAllCompanies(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.CompanyGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	context.JSON(http.StatusOK, pagination)
}

func (c *CompanyController) GetCompanyByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	claimsValue, _ := context.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	company, err := c.service.GetCompanyByID(uint(id), claims.UserID, claims.Role)
	if err != nil {
		context.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Company not found"})
		return
	}

	var companyResponse responses.CompanyGetResponse
	if err := mapstructure.Decode(company, &companyResponse); err != nil {
		context.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	context.JSON(http.StatusOK, companyResponse)
}

func (c *CompanyController) UpdateCompany(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := context.Get("input")
	companyRequest := validatedInput.(*requests.CompanyUpdateRequest)
	var company models.Company
	if err := mapstructure.Decode(companyRequest, &company); err != nil {
		context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	company.ID = uint(id)
	claimsValue, _ := context.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if companyRequest.Localiatation != nil {
		var localitationRequest models.Localitation
		if err := mapstructure.Decode(companyRequest.Localiatation, &localitationRequest); err != nil {
			context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
			return
		}

		localitation, err := c.localitationService.AssignLocalitation(&localitationRequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}

		company.LocalitationID = &localitation.ID
	}

	if err := c.service.UpdateCompany(&company, claims.UserID, claims.Role); err != nil {
		context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, company)
}

func (c *CompanyController) DeleteCompany(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := context.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := c.service.DeleteCompany(uint(id), claims.UserID, claims.Role); err != nil {
		context.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
