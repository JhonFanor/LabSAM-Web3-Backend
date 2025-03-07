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
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Image       string     `json:"image,omitempty"`
	Source      string     `json:"source,omitempty"`
	Link        string     `json:"link,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
}
