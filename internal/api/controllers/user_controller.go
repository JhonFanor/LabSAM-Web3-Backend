package controllers

import (
	"encoding/json"
	"fmt"
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

type UserControllerParams struct {
	fx.In
	UserService           services.UserService
	BusinessUserService   services.BusinessUserService
	RegularUserService    services.RegularUserService
	UniversityUserService services.UniversityUserService
	LocationService       services.LocationService
	ContactService        services.ContactService
	RoleService           services.RoleService
}

type UserController struct {
	service               services.UserService
	businessUserService   services.BusinessUserService
	regularUserService    services.RegularUserService
	universityUserService services.UniversityUserService
	locationService       services.LocationService
	contactService        services.ContactService
	roleService           services.RoleService
}

func NewUserController(p UserControllerParams) *UserController {
	return &UserController{
		service:               p.UserService,
		businessUserService:   p.BusinessUserService,
		regularUserService:    p.RegularUserService,
		universityUserService: p.UniversityUserService,
		locationService:       p.LocationService,
		contactService:        p.ContactService,
		roleService:           p.RoleService,
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

func (u *UserController) GetUserForProfile(c *gin.Context) {
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

func (u *UserController) UpdateRegularUser(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.RegularUserUpdateRequest)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user models.User
	if err := mapstructure.Decode(*input, &user); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	var regularUser models.RegularUser
	if err := mapstructure.Decode(*input, &regularUser); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	if input.LocationRequest != nil && input.LocationRequest.Country != "" && input.LocationRequest.Country != "" {
		var location models.Location
		if err := mapstructure.Decode(*input.LocationRequest, &location); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: consts.ErrorMapConst,
			})
			return
		}
		locationCreate, _ := u.locationService.CreateLocation(&location)
		regularUser.LocationID = &locationCreate.ID
	}

	if input.ContactRequest != nil && input.ContactRequest.Phone != "" && input.ContactRequest.Website != "" {
		var contact models.Contact
		if err := mapstructure.Decode(*input.ContactRequest, &contact); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: consts.ErrorMapConst,
			})
			return
		}

		u.contactService.UpdateContact(&contact)
	}

	user.ID = uint(id)
	if input.Password != nil {
		user.Password = *input.Password
	}
	regularUser.UserID = uint(id)

	err = u.service.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	err = u.regularUserService.UpdateRegularUser(&regularUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Regular user successfully updated",
	})
}

func (u *UserController) UpdateBusinessUser(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.BusinessUserUpdateRequest)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user models.User
	if err := mapstructure.Decode(*input, &user); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	var businessUser models.BusinessUser
	if err := mapstructure.Decode(*input, &businessUser); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	if input.LocationRequest != nil && input.LocationRequest.Country != "" && input.LocationRequest.Country != "" {
		var location models.Location
		if err := mapstructure.Decode(*input.LocationRequest, &location); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: consts.ErrorMapConst,
			})
			return
		}

		locationCreate, err := u.locationService.CreateLocation(&location)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		businessUser.LocationID = &locationCreate.ID
	}

	if input.ContactRequest != nil && input.ContactRequest.Phone != "" && input.ContactRequest.Website != "" {
		var contact models.Contact
		if err := mapstructure.Decode(*input.ContactRequest, &contact); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: consts.ErrorMapConst,
			})
			return
		}

		contactCreate, _ := u.contactService.CreateContact(&contact)

		businessUser.ContactID = &contactCreate.ID
	}

	if input.Password != nil {
		user.Password = *input.Password
	}

	user.ID = uint(id)
	businessUser.UserID = uint(id)
	fmt.Print(user)
	u.service.UpdateUser(&user)
	u.businessUserService.UpdateBusinessUser(&businessUser)

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Business user successfully updated",
	})
}

func (u *UserController) UpdateUniversityUser(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.UniversityUserUpdateRequest)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var user models.User
	if err := mapstructure.Decode(*input, &user); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	var universityUser models.UniversityUser
	if err := mapstructure.Decode(*input, &universityUser); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	if input.LocationRequest != nil && input.LocationRequest.Country != "" && input.LocationRequest.Country != "" {
		var location models.Location
		if err := mapstructure.Decode(*input.LocationRequest, &location); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: consts.ErrorMapConst,
			})
			return
		}
		locationCreate, _ := u.locationService.CreateLocation(&location)
		universityUser.LocationID = &locationCreate.ID
	}

	if input.ContactRequest != nil && input.ContactRequest.Phone != "" && input.ContactRequest.Website != "" {
		var contact models.Contact
		if err := mapstructure.Decode(*input.ContactRequest, &contact); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Error: consts.ErrorMapConst,
			})
			return
		}
		contactCreate, _ := u.contactService.CreateContact(&contact)

		universityUser.ContactID = &contactCreate.ID
	}
	if input.UniversityTypeRequest != nil {
		universityUser.UniversityTypeID = input.UniversityTypeRequest.ID
	}

	if input.Password != nil {
		user.Password = *input.Password
	}

	user.ID = uint(id)
	universityUser.UserID = uint(id)

	u.service.UpdateUser(&user)
	u.universityUserService.UpdateUniversityUser(&universityUser)

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "University user successfully updated",
	})
}
