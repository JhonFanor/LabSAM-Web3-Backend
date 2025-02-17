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

type InvestigationControllerParams struct {
	fx.In
	InvestigationService services.InvestigationService
}

type InvestigationController struct {
	service services.InvestigationService
}

func NewInvestigationController(p InvestigationControllerParams) *InvestigationController {
	return &InvestigationController{
		service: p.InvestigationService,
	}
}

// CreateInvestigation godoc
// @Summary Create a new Investigation entry
// @Description This endpoint creates a new Investigation record. Requires authentication.
// @Tags Investigation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.InvestigationRequest true "Investigation Information"
// @Success 201 {object} models.Investigation "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /investigation/create [post]
func (i *InvestigationController) CreateInvestigation(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	investigationRequest := validatedInput.(*requests.InvestigationRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var investigation models.Investigation
	if err := mapstructure.Decode(investigationRequest, &investigation); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdInvestigation, err := i.service.CreateInvestigation(&investigation, claims.UserID, investigationRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdInvestigation)
}

func (i *InvestigationController) GetAllInvestigations(c *gin.Context) {
	pagination, err := i.service.GetAllInvestigations(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

func (i *InvestigationController) GetInvestigationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	investigation, err := i.service.GetInvestigationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Investigation not found"})
		return
	}

	c.JSON(http.StatusOK, investigation)
}

func (i *InvestigationController) UpdateInvestigation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	investigationRequest := validatedInput.(*requests.InvestigationUpdateRequest)
	var investigation models.Investigation
	if err := mapstructure.Decode(investigationRequest, &investigation); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	investigation.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := i.service.UpdateInvestigation(&investigation, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, investigation)
}

func (i *InvestigationController) DeleteInvestigation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := i.service.DeleteInvestigation(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
