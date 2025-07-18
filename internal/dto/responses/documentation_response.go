package responses

type DocumentationGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type DocumentationGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type DocumentationGetResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Link        string                `json:"link"`
	IsApproved  *bool                 `json:"is_approved"`
	User        UserMinimalResponse   `json:"user"`
	Subtopics   []SubtopicGetResponse `json:"subtopics"`
}
