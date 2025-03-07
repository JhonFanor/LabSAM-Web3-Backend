package requests

type TopicRequest struct {
	Name string `json:"name" validate:"required"`
}
