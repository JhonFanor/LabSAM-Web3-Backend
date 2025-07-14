package customerrors

import "errors"

var (
	ErrInvalidData  = errors.New("invalid data provided")
	ErrInvalidID    = errors.New("invalid ID")
	ErrUnauthorized = errors.New("unauthorized action")
	ErrNoUpdates    = errors.New("no hay cambios para actualizar")
)
