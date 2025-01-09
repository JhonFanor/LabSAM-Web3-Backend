package controllers

import (
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type AuthControllerParams struct {
	fx.In
	AuthService services.AuthService
	Validator   *validator.Validate
}

type AuthController struct {
	AuthService services.AuthService
	Validator   *validator.Validate
}

func NewAuthController(p AuthControllerParams) *AuthController {
	return &AuthController{
		AuthService: p.AuthService,
		Validator:   p.Validator,
	}
}

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Register a new user with the provided details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body requests.UserRequest true "User registration data"
// @Success      200 {object} responses.UserResponse
// @Failure      400 {object} responses.ErrorResponse "Validation failed or bad request"
// @Failure      500 {object} responses.ErrorResponse "Internal server error"
// @Router       /auth/register [post]
func (a *AuthController) RegisterUser(c *gin.Context) {
	var input requests.UserRequest

	var user models.User
	if err := mapstructure.Decode(input, &user); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Error processing the request data",
		})
		return
	}

	createdUser, err := a.AuthService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, responses.UserResponse{
		Message: "User successfully registered",
		User:    *createdUser,
	})
}
