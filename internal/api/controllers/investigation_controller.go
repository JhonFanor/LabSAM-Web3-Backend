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

	investigation.Date = investigationRequest.Date

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

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.InvestigationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (i *InvestigationController) GetAllInvestigationsByUserID(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := i.service.GetAllInvestigationsByUserID(c, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.InvestigationGetAllByUserIDResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (i *InvestigationController) GetAllInvestigationsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := i.service.GetAllInvestigationsNotApproved(c, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.InvestigationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (i *InvestigationController) CountInvestigationsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	count, err := i.service.CountInvestigationsNotApproved(claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.CountResponse{Count: count})
}

func (i *InvestigationController) CountInvestigationBySubtopic(c *gin.Context) {
	count, err := i.service.CountInvestigationBySubtopic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, count)
}

func (i *InvestigationController) GetInvestigationByID(c *gin.Context) {
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

	investigation, err := i.service.GetInvestigationByID(uint(id), userID, role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Investigation not found"})
		return
	}

	var investigationResponse responses.InvestigationGetResponse
	if err := mapstructure.Decode(investigation, &investigationResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, investigationResponse)
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

	investigation.Date = investigationRequest.Date

	investigation.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := i.service.UpdateInvestigation(&investigation, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, investigation)
}

func (i *InvestigationController) SetInvestigationApproval(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	req := validatedInput.(*requests.ApprovalRequest)

	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	err = i.service.SetInvestigationApproval(uint(id), req.Approved, claims.UserID, claims.Role)
	if err != nil {
		if err == customerrors.ErrUnauthorized {
			c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Approval status updated"})
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
