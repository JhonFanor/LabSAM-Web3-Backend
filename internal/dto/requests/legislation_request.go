package requests

type LegislationRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}
