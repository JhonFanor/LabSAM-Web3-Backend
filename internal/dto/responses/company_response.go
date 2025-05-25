package responses

type CompanyGetAllResponse struct {
	ID   uint                `json:"id"`
	Name string              `json:"name"`
	User UserMinimalResponse `json:"user"`
}

type CompanyGetResponse struct {
	ID           uint                      `json:"id"`
	Name         string                    `json:"name"`
	Industry     string                    `json:"industry"`
	Website      string                    `json:"website"`
	Email        string                    `json:"email"`
	User         UserMinimalResponse       `json:"user"`
	Localitation LocalitationResponse      `json:"localitation"`
	Subtopics    []SubtopicMinimalResponse `json:"subtopics"`
}
