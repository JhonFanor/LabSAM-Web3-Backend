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

type DeniedPermissionUserControllerParams struct {
	fx.In
	DeniedPermissionUserService services.DeniedPermissionUserService
	PermissionService           services.PermissionService
}

type DeniedPermissionUserController struct {
	service           services.DeniedPermissionUserService
	permissionService services.PermissionService
}

func NewDeniedPermissionUserController(p DeniedPermissionUserControllerParams) *DeniedPermissionUserController {
	return &DeniedPermissionUserController{
		service:           p.DeniedPermissionUserService,
		permissionService: p.PermissionService,
	}
}

func (dc *DeniedPermissionUserController) GetAllDeniedPermissionsByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid user ID"})
		return
	}

	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	denied, err := dc.service.GetAllByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var permissionResponses []responses.PermissionResponse
	for _, pu := range denied {
		perm, err := dc.permissionService.GetPermissionByID(pu.PermissionID)
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

func (dc *DeniedPermissionUserController) AssignDeniedPermission(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	var req struct {
		PermissionID uint `json:"permission_id"`
		UserID       uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid request"})
		return
	}

	if err := dc.service.Assign(req.PermissionID, req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (dc *DeniedPermissionUserController) RevokeDeniedPermission(c *gin.Context) {
	claimsValue, _ := c.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	var req struct {
		PermissionID uint `json:"permission_id"`
		UserID       uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid request"})
		return
	}

	if err := dc.service.Revoke(req.PermissionID, req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
