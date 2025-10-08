package requests

type CurrencyTypeRequest struct {
	Code   string `json:"code" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Symbol string `json:"symbol" validate:"required"`
}
