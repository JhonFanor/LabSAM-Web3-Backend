package responses

import "time"

type LegislationGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	Logo  string              `json:"logo"`
	Date  *time.Time          `json:"date"`
	User  UserMinimalResponse `json:"user"`
}

type LegislationGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Logo       string              `json:"logo"`
	Date       *time.Time          `json:"date"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type LegislationGetResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Logo        string                `json:"logo"`
	Date        *time.Time            `json:"date"`
	Link        string                `json:"link"`
	User        UserMinimalResponse   `json:"user"`
	TypeOfLaw   TypeOfLawResponse     `json:"type_of_law"`
	IsApproved  *bool                 `json:"is_approved"`
	Subtopics   []SubtopicGetResponse `json:"subtopics"`
}
