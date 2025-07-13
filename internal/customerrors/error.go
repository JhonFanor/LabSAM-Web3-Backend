package customerrors

import "errors"

var (
	ErrInvalidData  = errors.New("invalid data provided")
	ErrInvalidID    = errors.New("invalid ID")
	ErrForbidden    = errors.New("no tienes permiso para acceder a este recurso")
	ErrUnauthorized = errors.New("unauthorized action")
	ErrNoUpdates    = errors.New("no hay cambios para actualizar")
)
