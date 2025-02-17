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

type LegislationControllerParams struct {
	fx.In
	LegislationService services.LegislationService
}

type LegislationController struct {
	service services.LegislationService
}

func NewLegislationController(p LegislationControllerParams) *LegislationController {
	return &LegislationController{
		service: p.LegislationService,
	}
}

// CreateLegislation godoc
// @Summary Create a new Legislation entry
// @Description This endpoint creates a new Legislation record. Requires authentication.
// @Tags Legislation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.LegislationRequest true "Legislation Information"
// @Success 201 {object} models.Legislation "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /legislation/create [post]
func (l *LegislationController) CreateLegislation(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	legislationRequest := validatedInput.(*requests.LegislationRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var legislation models.Legislation
	if err := mapstructure.Decode(legislationRequest, &legislation); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdLegislation, err := l.service.CreateLegislation(&legislation, claims.UserID, legislationRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdLegislation)
}

func (l *LegislationController) GetAllLegislations(c *gin.Context) {
	pagination, err := l.service.GetAllLegislations(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

func (l *LegislationController) GetLegislationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	legislation, err := l.service.GetLegislationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Legislation not found"})
		return
	}

	c.JSON(http.StatusOK, legislation)
}

func (l *LegislationController) UpdateLegislation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	legislationRequest := validatedInput.(*requests.LegislationUpdateRequest)
	var legislation models.Legislation
	if err := mapstructure.Decode(legislationRequest, &legislation); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	legislation.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := l.service.UpdateLegislation(&legislation, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, legislation)
}

func (l *LegislationController) DeleteLegislation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := l.service.DeleteLegislation(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
