package controllers

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PermissionRoleControllerParams struct {
	fx.In
	PermissionRoleService services.PermissionRoleService
}

type PermissionRoleController struct {
	service services.PermissionRoleService
}

func NewPermissionRoleController(p PermissionRoleControllerParams) *PermissionRoleController {
	return &PermissionRoleController{
		service: p.PermissionRoleService,
	}
}

func (pc *PermissionRoleController) GetAllPermissionsByRoleExcludingDenied(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}
	roleIDParam := c.Param("role_id")
	userIDParam := c.Param("user_id")

	roleID, err := strconv.Atoi(roleIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid role ID"})
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	permissions, err := pc.service.GetAllPermissionsByRoleExcludingDenied(uint(roleID), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var permissionResponses []responses.PermissionResponse
	for _, perm := range permissions {
		permissionResponses = append(permissionResponses, responses.PermissionResponse{
			ID:   perm.ID,
			Name: perm.Name,
		})
	}

	c.JSON(http.StatusOK, permissionResponses)
}
