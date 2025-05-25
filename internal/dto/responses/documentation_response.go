package responses

type DocumentationGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type DocumentationGetResponse struct {
	ID          uint                      `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Link        string                    `json:"link"`
	User        UserMinimalResponse       `json:"user"`
	Subtopics   []SubtopicMinimalResponse `json:"subtopics"`
}
