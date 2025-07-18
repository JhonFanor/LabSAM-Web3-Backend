package responses

type LegislationGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type LegislationGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Title      string              `json:"title"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type LegislationGetResponse struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Link        string                `json:"link"`
	User        UserMinimalResponse   `json:"user"`
	IsApproved  *bool                 `json:"is_approved"`
	Subtopics   []SubtopicGetResponse `json:"subtopics"`
}
