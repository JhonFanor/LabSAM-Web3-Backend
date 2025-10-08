package requests

import "time"

type LegislationRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Logo        string    `json:"logo"`
	Date        time.Time `json:"date" validate:"required"`
	Link        string    `json:"link" validate:"required"`
	TypeOfLawID uint      `json:"type_of_law_id" validate:"required"`
	SubtopicIDs []uint    `json:"subtopic_ids" validate:"required"`
}

type LegislationUpdateRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Logo        string    `json:"logo"`
	Date        time.Time `json:"date" validate:"required"`
	Link        string    `json:"link" validate:"required"`
	TypeOfLawID uint      `json:"type_of_law_id" validate:"required"`
}
