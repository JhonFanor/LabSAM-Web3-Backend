package responses

import (
	"time"
)

type NewsGetAllResponse struct {
	ID        uint                      `json:"id"`
	Title     string                    `json:"title"`
	Image     string                    `json:"image"`
	Date      *time.Time                `json:"date"`
	Subtopics []SubtopicMinimalResponse `json:"subtopics"`
	User      UserMinimalResponse       `json:"user"`
}

type NewGetResponse struct {
	ID          uint                      `json:"id,omitempty"`
	Title       string                    `json:"title,omitempty"`
	Description string                    `json:"description,omitempty"`
	Image       string                    `json:"image,omitempty"`
	Link        string                    `json:"link,omitempty"`
	Date        *time.Time                `json:"date,omitempty"`
	User        UserMinimalResponse       `json:"user,omitempty"`
	Subtopics   []SubtopicMinimalResponse `json:"subtopics,omitempty"`
}
