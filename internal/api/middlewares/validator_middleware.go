package middlewares

import (
	"lamsam-web3-backend/internal/dto/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidatorMiddleware struct {
	Validator *validator.Validate
}

func NewValidatorMiddleware(validator *validator.Validate) *ValidatorMiddleware {
	return &ValidatorMiddleware{
		Validator: validator,
	}
}

func (vm *ValidatorMiddleware) ValidateInput(input interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(input); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Error: "Invalid input data. Please check your request payload.",
			})
			c.Abort()
			return
		}

		if err := vm.Validator.Struct(input); err != nil {
			var validationErrors []string

			if validationErrs, ok := err.(validator.ValidationErrors); ok {
				for _, validationErr := range validationErrs {
					validationErrors = append(validationErrors, validationErr.Field()+" validation failed: "+validationErr.Tag())
				}
			}

			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Error:   "Validation failed",
				Details: validationErrors,
			})
			c.Abort()
			return
		}
		c.Set("input", input)
		c.Next()
	}
}
