package controllers

import (
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"

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
