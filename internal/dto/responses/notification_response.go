package responses

type NotificationGetAllResponse struct {
	ID           uint                 `json:"id"`
	Message      string               `json:"message"`
	ResourceType *string              `json:"resource_type"`
	ResourceID   *int                 `json:"resource_id"`
	Action       string               `json:"action"`
	Sender       *UserMinimalResponse `json:"sender"`
	IsRead       bool                 `json:"is_read"`
}

type NotificationGetResponse struct {
	ID           uint                 `json:"id"`
	Message      string               `json:"message"`
	ResourceType *string              `json:"resource_type"`
	ResourceID   *int                 `json:"resource_id"`
	Action       string               `json:"action"`
	Sender       *UserMinimalResponse `json:"sender"`
}
