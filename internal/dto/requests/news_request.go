package requests

import "time"

type NewsRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Source      string    `json:"source" validate:"required"`
	Link        string    `json:"link" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	SubtopicIDs []uint    `json:"subtopic_ids" validate:"required"`
}

type NewsUpdateRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Source      string    `json:"source"`
	Link        string    `json:"link"`
	Date        time.Time `json:"date"`
}
