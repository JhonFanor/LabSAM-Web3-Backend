package validations

import (
	"strings"

	"github.com/go-playground/validator/v10"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 60

func PasswordValidation(fl validator.FieldLevel) bool {
	password := strings.TrimSpace(fl.Field().String())
	err := passwordvalidator.Validate(password, minEntropyBits)
	if err != nil {
		return false
	}
	return true
}
