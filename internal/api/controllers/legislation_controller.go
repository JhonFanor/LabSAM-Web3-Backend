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

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.LegislationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (l *LegislationController) GetAllLegislationsByUserID(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := l.service.GetAllLegislationsByUserID(c, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.BankOfResumeGetAllByUserIDResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (l *LegislationController) GetAllLegislationsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := l.service.GetAllLegislationsNotApproved(c, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.LegislationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (l *LegislationController) CountLegislationsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	count, err := l.service.CountLegislationsNotApproved(claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.CountResponse{Count: count})
}

func (l *LegislationController) GetLegislationByID(c *gin.Context) {
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

	legislation, err := l.service.GetLegislationByID(uint(id), userID, role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Legislation not found"})
		return
	}

	var legislationResponse responses.LegislationGetResponse
	if err := mapstructure.Decode(legislation, &legislationResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, legislationResponse)
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
