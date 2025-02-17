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

// CreateDocumentation godoc
// @Summary Create a new Documentation entry
// @Description This endpoint creates a new Documentation record. Requires authentication.
// @Tags Documentation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.DocumentationRequest true "Documentation Information"
// @Success 201 {object} models.Documentation "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /documentation/create [post]
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
	c.JSON(http.StatusOK, pagination)
}

func (d *DocumentationController) GetDocumentationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	documentation, err := d.service.GetDocumentationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Documentation not found"})
		return
	}

	c.JSON(http.StatusOK, documentation)
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
