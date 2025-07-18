package responses

import "time"

type EducationalOfferGetAllResponse struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Institution string              `json:"institution"`
	StartDate   *time.Time          `json:"start_date"`
	EndDate     *time.Time          `json:"end_date"`
	Cost        float64             `json:"cost"`
	User        UserMinimalResponse `json:"user"`
}

type EducationalOfferGetAllByUserIDResponse struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Institution string              `json:"institution"`
	StartDate   *time.Time          `json:"start_date"`
	EndDate     *time.Time          `json:"end_date"`
	Cost        float64             `json:"cost"`
	IsApproved  *bool               `json:"is_approved"`
	User        UserMinimalResponse `json:"user"`
}

type EducationalOfferGetResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Institution string                `json:"institution"`
	StartDate   *time.Time            `json:"start_date"`
	EndDate     *time.Time            `json:"end_date"`
	Cost        float64               `json:"cost"`
	Description string                `json:"description"`
	Link        string                `json:"link"`
	IsApproved  *bool                 `json:"is_approved"`
	User        UserMinimalResponse   `json:"user"`
	Subtopics   []SubtopicGetResponse `json:"subtopics"`
}
