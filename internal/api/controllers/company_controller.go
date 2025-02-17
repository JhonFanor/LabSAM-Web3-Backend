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

type CompanyControllerParams struct {
	fx.In
	CompanyService services.CompanyService
}

type CompanyController struct {
	service services.CompanyService
}

func NewCompanyController(p CompanyControllerParams) *CompanyController {
	return &CompanyController{
		service: p.CompanyService,
	}
}

// CreateCompany godoc
// @Summary Create a new Company entry
// @Description This endpoint creates a new Company record. Requires authentication.
// @Tags Company
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.CompanyRequest true "Company Information"
// @Success 201 {object} models.Company "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /company/create [post]
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
	context.JSON(http.StatusOK, pagination)
}

func (c *CompanyController) GetCompanyByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	company, err := c.service.GetCompanyByID(uint(id))
	if err != nil {
		context.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Company not found"})
		return
	}

	context.JSON(http.StatusOK, company)
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
