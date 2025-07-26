package controllers

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PermissionControllerParams struct {
	fx.In
	PermissionService services.PermissionService
	RoleService       services.RoleService
}

type PermissionController struct {
	service services.PermissionService
}

func NewPermissionController(p PermissionControllerParams) *PermissionController {
	return &PermissionController{
		service: p.PermissionService,
	}
}

func (pc *PermissionController) GetAllAssignablePermissionsToUser(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}
	validatedInput, _ := c.Get("input")
	permissionRequest := validatedInput.(*requests.PermissionRequest)

	permissions, err := pc.service.GetAllAssignablePermissionsToUser(permissionRequest.UserID, permissionRequest.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, permissions)
}
