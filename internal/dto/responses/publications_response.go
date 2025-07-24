package responses

type PublicationsResponse struct {
	ResourceType string      `json:"resource_type"`
	Data         interface{} `json:"data"`
}
