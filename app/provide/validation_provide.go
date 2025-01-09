package provide

import (
	"lamsam-web3-backend/internal/validations"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func NewValidatorRegister() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("passwordValidation", validations.PasswordValidation)
	return v
}

func ValidatorProvide() fx.Option {
	return fx.Options(
		fx.Provide(
			NewValidatorRegister,
		),
	)
}
