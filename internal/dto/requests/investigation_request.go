package requests

import "time"

type InvestigationRequest struct {
	Title       string     `json:"title" validate:"required"`
	Author      string     `json:"author" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Logo        string     `json:"logo"`
	Date        *time.Time `json:"date" validate:"required"`
	Link        string     `json:"link" validate:"required"`
	SubtopicIDs []uint     `json:"subtopic_ids" validate:"required"`
}

type InvestigationUpdateRequest struct {
	Title       string     `json:"title" validate:"required"`
	Author      string     `json:"author" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Logo        string     `json:"logo"`
	Date        *time.Time `json:"date" validate:"required"`
	Link        string     `json:"link" validate:"required"`
}
