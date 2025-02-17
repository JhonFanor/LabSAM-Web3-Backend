package requests

type JobExchangeRequest struct {
	Title       string `json:"title" validate:"required"`
	Company     string `json:"company" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"`
	SalaryRange string `json:"salary_range" validate:"required"`
	Status      string `json:"status" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type JobExchangeUpdateRequest struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Type        string `json:"type"`
	SalaryRange string `json:"salary_range"`
	Status      string `json:"status"`
}
