package controllers

import (
	"encoding/json"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

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
		ID:          user.ID,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Role:        role.Name,
		Permissions: claims.Permissions,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	pagination, err := u.service.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.UserGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (u *UserController) GetUserForAdminByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	user, err := u.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "User not found"})
		return
	}

	var userResponse responses.UserGetResponse
	jsonData, _ := json.Marshal(user)
	_ = json.Unmarshal(jsonData, &userResponse)

	userResponse.RoleID = user.RoleID

	c.JSON(http.StatusOK, userResponse)
}
