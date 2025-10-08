package requests

type DocumentationRequest struct {
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required"`
	SubtopicIDs []uint `json:"subtopic_ids" validate:"required"`
}

type DocumentationUpdateRequest struct {
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required"`
}
