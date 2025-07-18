package responses

type JobBoardGetAllResponse struct {
	ID      uint                `json:"id"`
	Title   string              `json:"title"`
	Company string              `json:"company"`
	User    UserMinimalResponse `json:"user"`
}

type JobBoardGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	Company    string              `json:"company"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type JobBoardGetResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Company     string                `json:"company"`
	Description string                `json:"description"`
	Type        string                `json:"type"`
	SalaryRange string                `json:"salary_range"`
	Link        string                `json:"link"`
	IsApproved  *bool                 `json:"is_approved"`
	User        UserMinimalResponse   `json:"user"`
	Subtopics   []SubtopicGetResponse `json:"subtopics"`
}
