package responses

import "time"

type InvestigationGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type InvestigationGetResponse struct {
	ID          uint                      `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Date        time.Time                 `json:"date"`
	Link        string                    `json:"link"`
	User        UserMinimalResponse       `json:"user"`
	Subtopics   []SubtopicMinimalResponse `json:"subtopics"`
}
