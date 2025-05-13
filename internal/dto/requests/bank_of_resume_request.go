package requests

type BankOfResumeCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Photo       string `json:"photo" validate:"required"`
	Summary     string `json:"summary" validate:"required"`
	Link        string `json:"link" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type BankOfResumeUpdateRequest struct {
	Title   string `json:"title"`
	Photo   string `json:"photo"`
	Summary string `json:"summary"`
	Link    string `json:"link"`
}
