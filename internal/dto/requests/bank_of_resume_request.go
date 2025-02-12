package requests

type BankOfResumeRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Skills      string `json:"skills" validate:"required"`
	Experience  string `json:"experience" validate:"required"`
	Education   string `json:"education" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}
