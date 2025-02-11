package customerrors

import "errors"

var (
	ErrInvalidData = errors.New("invalid data provided")
	ErrInvalidID   = errors.New("invalid ID")
)
