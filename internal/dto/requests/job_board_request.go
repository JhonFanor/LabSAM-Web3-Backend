package requests

type JobBoardRequest struct {
	Title       string `json:"title" validate:"required"`
	Company     string `json:"company" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"`
	SalaryRange string `json:"salary_range" validate:"required"`
	Link        string `json:"link" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type JobBoardUpdateRequest struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Type        string `json:"type"`
	SalaryRange string `json:"salary_range"`
	Link        string `json:"link"`
}
