package requests

import "time"

type EducationalOfferRequest struct {
	Title       string     `json:"title" validate:"required"`
	Institution string     `json:"institution" validate:"required"`
	StartDate   *time.Time `json:"start_date" validate:"required"`
	EndDate     *time.Time `json:"end_date" validate:"required"`
	Cost        float64    `json:"cost"`
	Description string     `json:"description" validate:"required"`
	Link        string     `json:"link" validate:"required"`
	SubtopicIDs []uint     `json:"subtopic_ids" validate:"required"`
}

type EducationalUpdateOfferRequest struct {
	Title       string     `json:"title"`
	Institution string     `json:"institution"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Cost        float64    `json:"cost"`
	Description string     `json:"description"`
	Website     string     `json:"website"`
}
