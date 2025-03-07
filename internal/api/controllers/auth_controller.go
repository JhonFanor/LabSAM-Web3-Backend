package controllers

import (
	"lamsam-web3-backend/config"
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/internal/utils"
	"lamsam-web3-backend/internal/validations"
	"lamsam-web3-backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type AuthControllerParams struct {
	fx.In
	Validator         *validator.Validate
	JwtConfig         *config.JwtConfig
	AuthService       services.AuthService
	UserService       services.UserService
	RoleService       services.RoleService
	PermissionService services.PermissionService
}

type AuthController struct {
	Validator         *validator.Validate
	JwtConfig         *config.JwtConfig
	AuthService       services.AuthService
	UserService       services.UserService
	RoleService       services.RoleService
	PermissionService services.PermissionService
}

func NewAuthController(p AuthControllerParams) *AuthController {
	return &AuthController{
		Validator:         p.Validator,
		JwtConfig:         p.JwtConfig,
		AuthService:       p.AuthService,
		UserService:       p.UserService,
		RoleService:       p.RoleService,
		PermissionService: p.PermissionService,
	}
}

// RegisterBusinessUser godoc
// @Summary Register a new business user
// @Description This API endpoint registers a new business user with all the necessary details.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param input body requests.BusinessUserRequest true "Business User Information"
// @Success 200 {object} responses.UserResponse "Business user successfully registered"
// @Failure 400 {object} responses.ErrorResponse "Error registering business user"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /auth/register/business [post]
func (a *AuthController) RegisterBusinessUser(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.BusinessUserRequest)

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

	createdUser, err := a.AuthService.RegisterBusinessUser(&user, &businessUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Business user successfully registered",
		Data:    *createdUser,
	})
}

// RegisterRegularUser godoc
// @Summary Register a new regular user
// @Description This API endpoint registers a new regular user with all the necessary details.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param input body requests.RegularUserRequest true "Regular User Information"
// @Success 200 {object} responses.UserResponse "Regular user successfully registered"
// @Failure 400 {object} responses.ErrorResponse "Error registering regular user"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /auth/register/regular [post]
func (a *AuthController) RegisterRegularUser(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.RegularUserRequest)

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

	createdUser, err := a.AuthService.RegisterRegularUser(&user, &regularUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Regular user successfully registered",
		Data:    *createdUser,
	})
}

// RegisterUniversityUser godoc
// @Summary Register a new university user
// @Description This API endpoint registers a new university user with all the necessary details.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param input body requests.UniversityUserRequest true "University User Information"
// @Success 200 {object} responses.UserResponse "University user successfully registered"
// @Failure 400 {object} responses.ErrorResponse "Error registering university user"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /auth/register/university [post]
func (a *AuthController) RegisterUniversityUser(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.UniversityUserRequest)

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

	createdUser, err := a.AuthService.RegisterUniversityUser(&user, &universityUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "University user successfully registered",
		Data:    *createdUser,
	})
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticate a user and generate access and refresh tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body requests.LoginRequest true "User login data"
// @Success      200 {object} responses.AuthResponse
// @Failure      400 {object} responses.ErrorResponse "Bad request, invalid credentials"
// @Failure      401 {object} responses.ErrorResponse "Invalid credentials"
// @Failure      500 {object} responses.ErrorResponse "Internal server error"
// @Router       /auth/login [post]
func (a *AuthController) Login(c *gin.Context) {
	validatedInput, _ := c.Get("input")

	input := validatedInput.(*requests.LoginRequest)

	var user *models.User
	var err error

	if utils.IsValidEmail(input.UsernameOrEmail) {
		user, err = a.UserService.FindUserByEmail(input.UsernameOrEmail)
	} else {
		user, err = a.UserService.FindUserByUsername(input.UsernameOrEmail)
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "Invalid credentials",
		})
		return
	}

	if !user.VerifyPassword(input.Password) {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "Invalid password",
		})
		return
	}

	role, _ := a.RoleService.GetRoleByID(user.RoleID)
	permissions, _ := a.PermissionService.GetAllPermissionsByUser(user.ID, user.RoleID)

	accessToken, err := security.GenerateAccessToken(user.ID, user.Username, user.Email, *role, permissions, a.JwtConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Could not generate access token",
		})
		return
	}

	refreshToken, err := security.GenerateRefreshToken(user.Username, a.JwtConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Could not generate refresh token",
		})
		return
	}

	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "localhost", false, true)

	c.JSON(http.StatusOK, responses.AuthResponse{
		AccessToken: accessToken,
	})
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Generate a new access token using a valid refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 {object} responses.AuthResponse
// @Failure      400 {object} responses.ErrorResponse "Bad request, invalid refresh token"
// @Failure      401 {object} responses.ErrorResponse "Invalid or expired refresh token"
// @Failure      500 {object} responses.ErrorResponse "Internal server error"
// @Router       /auth/token/refresh [post]
func (a *AuthController) RefreshToken(c *gin.Context) {
	// Obtener el refresh_token de la cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "No refresh token found",
		})
		return
	}

	claims, err := validations.ValidateRefreshToken(refreshToken, a.JwtConfig)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "Invalid or expired refresh token",
		})
		return
	}

	user, err := a.UserService.FindUserByUsername(claims.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Error: "User not found",
		})
		return
	}

	role, _ := a.RoleService.GetRoleByID(user.RoleID)
	permissions, _ := a.PermissionService.GetAllPermissionsByUser(user.ID, user.RoleID)

	newAccessToken, err := security.GenerateAccessToken(user.ID, user.Username, user.Email, *role, permissions, a.JwtConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Could not generate new access token",
		})
		return
	}

	c.JSON(http.StatusOK, responses.AuthResponse{
		AccessToken: newAccessToken,
	})
}

func (a *AuthController) Logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Sesión cerrada correctamente"})
}
