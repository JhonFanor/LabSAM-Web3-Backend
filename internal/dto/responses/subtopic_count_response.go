package responses

type SubtopicCountResponse struct {
	SubtopicName string `json:"subtopic_name"`
	Count        int64  `json:"count"`
}
