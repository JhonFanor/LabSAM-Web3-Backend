package requests

type RejectionCommentCreateRequest struct {
	ResourceType string `json:"resource_type" validate:"required"`
	ResourceID   uint   `json:"resource_id" validate:"required"`
	Comment      string `json:"comment" validate:"required"`
}

type RejectionCommentUpdateRequest struct {
	Comment string `json:"comment" validate:"required"`
}
