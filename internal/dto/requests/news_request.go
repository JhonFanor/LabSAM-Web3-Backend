package requests

import "time"

type NewsRequest struct {
	Title       string    `json:"title" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Link        string    `json:"link"`
	Date        time.Time `json:"date" validate:"required"`
	SubtopicIDs []uint    `json:"subtopic_ids" validate:"required"`
}

type NewsUpdateRequest struct {
	Title       string     `json:"title,omitempty"`
	Image       string     `json:"image,omitempty"`
	Description string     `json:"description,omitempty"`
	Link        string     `json:"link,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
}
