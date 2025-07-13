package responses

type SubtopicGetResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SubtopicMinimalResponse struct {
	Name string `json:"name"`
}
