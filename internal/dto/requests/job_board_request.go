package requests

import "time"

type JobBoardRequest struct {
	Title        string              `json:"title" validate:"required"`
	Logo         string              `json:"logo"`
	Company      string              `json:"company" validate:"required"`
	Description  string              `json:"description" validate:"required"`
	Type         string              `json:"type"`
	SalaryRange  string              `json:"salary_range"`
	Link         string              `json:"link" validate:"required"`
	StartDate    time.Time           `json:"start_date" validate:"required"`
	EndDate      time.Time           `json:"end_date"`
	CurrencyType CurrencyTypeRequest `json:"currency_type" validate:"required"`
	SubtopicIDs  []uint              `json:"subtopic_ids" validate:"required"`
}

type JobBoardUpdateRequest struct {
	Title        string              `json:"title" validate:"required"`
	Logo         string              `json:"logo"`
	Company      string              `json:"company" validate:"required"`
	Description  string              `json:"description" validate:"required"`
	Type         string              `json:"type"`
	SalaryRange  string              `json:"salary_range"`
	Link         string              `json:"link" validate:"required"`
	StartDate    time.Time           `json:"start_date" validate:"required"`
	EndDate      time.Time           `json:"end_date"`
	CurrencyType CurrencyTypeRequest `json:"currency_type" validate:"required"`
}
