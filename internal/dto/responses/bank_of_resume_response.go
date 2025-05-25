package responses

type BankOfResumeGetAllResponse struct {
	ID    uint                `json:"id"`
	Photo string              `json:"photo"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type BankOfResumeGetResponse struct {
	ID        uint                      `json:"id,omitempty"`
	Photo     string                    `json:"photo"`
	Title     string                    `json:"title"`
	Summary   string                    `json:"summary"`
	Link      string                    `json:"link"`
	User      UserMinimalResponse       `json:"user,omitempty"`
	Subtopics []SubtopicMinimalResponse `json:"subtopics,omitempty"`
}
