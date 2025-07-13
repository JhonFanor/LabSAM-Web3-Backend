package responses

type TopicGetAllResponse struct {
	ID        uint                  `json:"id"`
	Name      string                `json:"name"`
	Subtopics []SubtopicGetResponse `json:"subtopics"`
}
