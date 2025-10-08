package responses

import "time"

type JobBoardGetAllResponse struct {
	ID      uint                `json:"id"`
	Title   string              `json:"title"`
	Logo    string              `json:"logo"`
	Company string              `json:"company"`
	User    UserMinimalResponse `json:"user"`
}

type JobBoardGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Logo       string              `json:"logo"`
	Company    string              `json:"company"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type JobBoardGetResponse struct {
	ID           uint                  `json:"id"`
	Title        string                `json:"title"`
	Logo         string                `json:"logo"`
	Company      string                `json:"company"`
	Description  string                `json:"description"`
	Type         string                `json:"type"`
	SalaryRange  string                `json:"salary_range"`
	Link         string                `json:"link"`
	StartDate    time.Time             `json:"start_date"`
	EndDate      time.Time             `json:"end_date"`
	IsApproved   *bool                 `json:"is_approved"`
	User         UserMinimalResponse   `json:"user"`
	CurrencyType CurrencyTypeRespone   `json:"currency_type"`
	Subtopics    []SubtopicGetResponse `json:"subtopics"`
}
