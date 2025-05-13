package requests

type LegislationRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type LegislationUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
