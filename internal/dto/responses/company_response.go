package responses

type CompanyGetAllResponse struct {
	ID   uint                `json:"id"`
	Name string              `json:"name"`
	Logo string              `json:"logo"`
	User UserMinimalResponse `json:"user"`
}

type CompanyGetAllByUserIDResponse struct {
	ID         uint                `json:"id"`
	Name       string              `json:"name"`
	Logo       string              `json:"logo"`
	IsApproved *bool               `json:"is_approved"`
	User       UserMinimalResponse `json:"user"`
}

type CompanyGetResponse struct {
	ID           uint                  `json:"id"`
	Name         string                `json:"name"`
	Logo         string                `json:"logo"`
	Industry     string                `json:"industry"`
	Website      string                `json:"website"`
	Email        string                `json:"email"`
	Projects     string                `json:"projects"`
	IsApproved   *bool                 `json:"is_approved"`
	User         UserMinimalResponse   `json:"user"`
	Localitation *LocalitationResponse `json:"localitation"`
	Subtopics    []SubtopicGetResponse `json:"subtopics"`
}
