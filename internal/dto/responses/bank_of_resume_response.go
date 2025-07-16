package responses

type BankOfResumeGetAllResponse struct {
	ID    uint                `json:"id"`
	Photo string              `json:"photo"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type BankOfResumeGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Photo      string              `json:"photo"`
	Title      string              `json:"title"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type BankOfResumeGetResponse struct {
	ID        uint                      `json:"id"`
	Photo     string                    `json:"photo"`
	Title     string                    `json:"title"`
	Summary   string                    `json:"summary"`
	Link      string                    `json:"link"`
	User      UserMinimalResponse       `json:"user"`
	Subtopics []SubtopicMinimalResponse `json:"subtopics"`
}
