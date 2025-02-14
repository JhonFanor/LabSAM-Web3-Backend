package requests

type BankOfResumeCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Skills      string `json:"skills" validate:"required"`
	Experience  string `json:"experience" validate:"required"`
	Education   string `json:"education" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type BankOfResumeUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
	Experience  string `json:"experience"`
	Education   string `json:"education"`
}
