package requests

import "time"

type EventRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Link        string    `json:"link" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	SubtopicIDs []uint    `json:"subtopic_ids" validate:"required"`
}
