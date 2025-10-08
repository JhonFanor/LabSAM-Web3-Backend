package requests

import "time"

type EventRequest struct {
	Title            string               `json:"title" validate:"required"`
	Image            string               `json:"image"`
	Poster           string               `json:"poster"`
	Description      string               `json:"description" validate:"required"`
	Link             string               `json:"link"`
	RegistrationLink string               `json:"registration_link"`
	Date             *time.Time           `json:"date" validate:"required"`
	Localiatation    *LocalitationRequest `json:"localitation" validate:"required"`
	SubtopicIDs      []uint               `json:"subtopic_ids" validate:"required"`
}

type EventUpdateRequest struct {
	Title            string               `json:"title" validate:"required"`
	Image            string               `json:"image"`
	Poster           string               `json:"poster"`
	Description      string               `json:"description" validate:"required"`
	Link             string               `json:"link"`
	RegistrationLink string               `json:"registration_link"`
	Date             *time.Time           `json:"date" validate:"required"`
	Localiatation    *LocalitationRequest `json:"localitation" validate:"required"`
}
