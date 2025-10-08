package responses

import "time"

type InvestigationGetAllResponse struct {
	ID     uint                `json:"id"`
	Title  string              `json:"title"`
	Author string              `json:"author"`
	Logo   string              `json:"logo"`
	User   UserMinimalResponse `json:"user"`
}

type InvestigationGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Author     string              `json:"author"`
	Logo       string              `json:"logo"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type InvestigationGetResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Author      string                `json:"author"`
	Description string                `json:"description"`
	Logo        string                `json:"logo"`
	Date        *time.Time            `json:"date"`
	Link        string                `json:"link"`
	IsApproved  *bool                 `json:"is_approved"`
	User        UserMinimalResponse   `json:"user"`
	Subtopics   []SubtopicGetResponse `json:"subtopics"`
}
