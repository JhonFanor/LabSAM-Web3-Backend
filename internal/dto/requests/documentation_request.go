package requests

type DocumentationRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type DocumentationUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
