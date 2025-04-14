package requests

import "time"

type EventRequest struct {
	Title         string               `json:"title" validate:"required"`
	Description   string               `json:"description" validate:"required"`
	Link          string               `json:"link" validate:"required"`
	Date          time.Time            `json:"date" validate:"required"`
	Localiatation *LocalitationRequest `json:"localitation"`
	SubtopicIDs   []uint               `json:"subtopic_ids" validate:"required"`
}

type EventUpdateRequest struct {
	Title         string               `json:"title"`
	Description   string               `json:"description"`
	Link          string               `json:"link"`
	Date          time.Time            `json:"date"`
	Localiatation *LocalitationRequest `json:"localitation"`
}
