package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	re := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*(),.?":{}|<>]).{8,}$`
	matched, _ := regexp.MatchString(re, password)
	return matched
}
