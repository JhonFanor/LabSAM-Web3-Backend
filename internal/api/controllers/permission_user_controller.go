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

type PermissionUserControllerParams struct {
	fx.In
	PermissionUserService services.PermissionUserService
	PermissionService     services.PermissionService
}

type PermissionUserController struct {
	service           services.PermissionUserService
	permissionService services.PermissionService
}

func NewPermissionUserController(p PermissionUserControllerParams) *PermissionUserController {
	return &PermissionUserController{
		service:           p.PermissionUserService,
		permissionService: p.PermissionService,
	}
}

func (pc *PermissionUserController) AssignPermissionToUser(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	permissionUserRequest := validatedInput.(*requests.PermissionUserRequest)

	var permissionUser models.PermissionUser
	if err := mapstructure.Decode(permissionUserRequest, &permissionUser); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	if err := pc.service.AssignPermissionToUser(&permissionUser); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Permission assigned successfully"})
}

func (pc *PermissionUserController) RevokePermissionFromUser(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}
	permissionIDParam := c.Param("permission_id")
	userIDParam := c.Param("user_id")

	permissionID, err := strconv.Atoi(permissionIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid permission ID"})
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	if err := pc.service.RevokePermissionFromUser(uint(permissionID), uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Permission revoked successfully"})
}

func (pc *PermissionUserController) GetAllPermissionsByUser(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	permissionUsers, err := pc.service.GetAllPermissionsByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var permissionResponses []responses.PermissionResponse
	for _, pu := range permissionUsers {
		perm, err := pc.permissionService.GetPermissionByID(pu.PermissionID)
		if err != nil {
			continue
		}
		permissionResponses = append(permissionResponses, responses.PermissionResponse{
			ID:   perm.ID,
			Name: perm.Name,
		})
	}

	c.JSON(http.StatusOK, permissionResponses)
}
