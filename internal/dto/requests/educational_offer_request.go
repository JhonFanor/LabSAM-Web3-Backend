package requests

import "time"

type EducationalOfferRequest struct {
	Title           string               `json:"title" validate:"required"`
	Institution     string               `json:"institution" validate:"required"`
	Logo            string               `json:"logo"`
	StartDate       *time.Time           `json:"start_date" validate:"required"`
	EndDate         *time.Time           `json:"end_date" validate:"required"`
	Cost            float64              `json:"cost"`
	Description     string               `json:"description" validate:"required"`
	Link            string               `json:"link" validate:"required"`
	TypeEducationID uint                 `json:"type_education_id"`
	CurrencyType    *CurrencyTypeRequest `json:"currency_type"`
	SubtopicIDs     []uint               `json:"subtopic_ids" validate:"required"`
}

type EducationalUpdateOfferRequest struct {
	Title           string               `json:"title" validate:"required"`
	Institution     string               `json:"institution" validate:"required"`
	Logo            string               `json:"logo"`
	StartDate       *time.Time           `json:"start_date" validate:"required"`
	EndDate         *time.Time           `json:"end_date" validate:"required"`
	Cost            float64              `json:"cost"`
	Description     string               `json:"description" validate:"required"`
	Link            string               `json:"link" validate:"required"`
	TypeEducationID uint                 `json:"type_education_id"`
	CurrencyType    *CurrencyTypeRequest `json:"currency_type" validate:"required"`
}
