package models

type CurrencyType struct {
	ID     uint   `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (CurrencyType) TableName() string {
	return "currency_type"
}
