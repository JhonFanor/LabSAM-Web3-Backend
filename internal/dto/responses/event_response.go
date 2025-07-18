package responses

import "time"

type EventGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	Image string              `json:"image"`
	Date  time.Time           `json:"date"`
	User  UserMinimalResponse `json:"user"`
}

type EventGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Image      string              `json:"image"`
	Date       time.Time           `json:"date"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type EventGetResponse struct {
	ID           uint                  `json:"id"`
	Title        string                `json:"title"`
	Image        string                `json:"image"`
	Description  string                `json:"description"`
	Link         string                `json:"link"`
	Date         time.Time             `json:"date"`
	User         UserMinimalResponse   `json:"user"`
	IsApproved   *bool                 `json:"is_approved"`
	Localitation *LocalitationResponse `json:"localitation,omitempty"`
	Subtopics    []SubtopicGetResponse `json:"subtopics"`
}
