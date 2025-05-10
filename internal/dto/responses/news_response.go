package responses

import (
	"time"
)

type NewsGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	Image string              `json:"image"`
	Date  *time.Time          `json:"date"`
	User  UserMinimalResponse `json:"user"`
}

type NewGetResponse struct {
	ID          uint                      `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Link        string                    `json:"link"`
	Date        *time.Time                `json:"date"`
	User        UserMinimalResponse       `json:"user"`
	Subtopics   []SubtopicMinimalResponse `json:"subtopics"`
}
