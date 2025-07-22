package validations

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func PasswordValidation(fl validator.FieldLevel) bool {
	password := strings.TrimSpace(fl.Field().String())

	if len(password) < 8 {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	hasDigit := regexp.MustCompile(`\d`).MatchString
	hasSpecial := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString

	return hasLower(password) && hasUpper(password) && hasDigit(password) && hasSpecial(password)
}
