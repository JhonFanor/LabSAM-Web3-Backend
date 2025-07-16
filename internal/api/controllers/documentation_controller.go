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

type DocumentationControllerParams struct {
	fx.In
	DocumentationService services.DocumentationService
}

type DocumentationController struct {
	service services.DocumentationService
}

func NewDocumentationController(p DocumentationControllerParams) *DocumentationController {
	return &DocumentationController{
		service: p.DocumentationService,
	}
}

func (d *DocumentationController) CreateDocumentation(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	documentationRequest := validatedInput.(*requests.DocumentationRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var documentation models.Documentation
	if err := mapstructure.Decode(documentationRequest, &documentation); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdDocumentation, err := d.service.CreateDocumentation(&documentation, claims.UserID, documentationRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdDocumentation)
}

func (d *DocumentationController) GetAllDocumentations(c *gin.Context) {
	pagination, err := d.service.GetAllDocumentations(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.DocumentationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (d *DocumentationController) GetAllDocumentationsByUserID(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := d.service.GetAllDocumentationsByUserID(c, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.DocumentationGetAllByUserIDResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (d *DocumentationController) GetAllDocumentationsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := d.service.GetAllDocumentationsNotApproved(c, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.DocumentationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (d *DocumentationController) CountDocumentationsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	count, err := d.service.CountDocumentationsNotApproved(claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.CountResponse{Count: count})
}

func (d *DocumentationController) GetDocumentationByID(c *gin.Context) {
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

	documentation, err := d.service.GetDocumentationByID(uint(id), userID, role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Documentation not found"})
		return
	}

	var documentationResponse responses.DocumentationGetResponse
	if err := mapstructure.Decode(documentation, &documentationResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, documentationResponse)
}

func (d *DocumentationController) UpdateDocumentation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	documentationRequest := validatedInput.(*requests.DocumentationUpdateRequest)
	var documentation models.Documentation
	if err := mapstructure.Decode(documentationRequest, &documentation); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	documentation.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := d.service.UpdateDocumentation(&documentation, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, documentation)
}

func (d *DocumentationController) DeleteDocumentation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := d.service.DeleteDocumentation(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
