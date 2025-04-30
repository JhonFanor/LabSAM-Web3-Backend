package controllers

import (
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UserControllerParams struct {
	fx.In
	UserService services.UserService
	RoleService services.RoleService
}

type UserController struct {
	service     services.UserService
	roleService services.RoleService
}

func NewUserController(p UserControllerParams) *UserController {
	return &UserController{
		service:     p.UserService,
		roleService: p.RoleService,
	}
}

func (u *UserController) GetUserByID(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)
	user, err := u.service.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "User not found"})
		return
	}

	role, err := u.roleService.GetRoleByID(user.RoleID)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Role not found"})
		return
	}

	userResponse := responses.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Avatar: user.Avatar,
		Role:   role.Name,
	}

	c.JSON(http.StatusOK, userResponse)
}
