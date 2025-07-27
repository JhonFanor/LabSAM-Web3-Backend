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

type RejectionCommentControllerParams struct {
	fx.In
	RejectionCommentService services.RejectionCommentService
}

type RejectionCommentController struct {
	service services.RejectionCommentService
}

func NewRejectionCommentController(p RejectionCommentControllerParams) *RejectionCommentController {
	return &RejectionCommentController{
		service: p.RejectionCommentService,
	}
}

func (rc *RejectionCommentController) Create(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
	}

	validatedInput, _ := c.Get("input")
	rejectionCommentRequest := validatedInput.(*requests.RejectionCommentCreateRequest)
	var rejectionComment models.RejectionComment
	if err := mapstructure.Decode(rejectionCommentRequest, &rejectionComment); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	created, err := rc.service.CreateRejectionComment(&rejectionComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (rc *RejectionCommentController) Update(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	rejectionCommentRequest := validatedInput.(*requests.RejectionCommentUpdateRequest)
	var rejectionComment models.RejectionComment
	if err := mapstructure.Decode(rejectionCommentRequest, &rejectionComment); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	rejectionComment.ID = uint(id)

	if err := rc.service.UpdateRejectionComment(&rejectionComment); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Rejection comment updated"})
}

func (rc *RejectionCommentController) Delete(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	if err := rc.service.DeleteRejectionComment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Rejection comment deleted"})
}

func (rc *RejectionCommentController) GetAllByResource(c *gin.Context) {
	resourceType := c.Param("resource_type")
	resourceID, err := strconv.Atoi(c.Param("resource_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	comments, err := rc.service.GetAllByResource(resourceType, uint(resourceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var response []responses.RejectionCommentResponse
	for _, cmt := range comments {
		response = append(response, responses.RejectionCommentResponse{
			ID:      cmt.ID,
			Comment: cmt.Comment,
		})
	}

	c.JSON(http.StatusOK, response)
}
