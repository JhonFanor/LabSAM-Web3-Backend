package requests

type CompanyRequest struct {
	Name          string               `json:"name" validate:"required"`
	Logo          string               `json:"logo"`
	Industry      string               `json:"industry" validate:"required"`
	Website       string               `json:"website"`
	Email         string               `json:"email"`
	Projects      string               `json:"projects"`
	Localiatation *LocalitationRequest `json:"localitation"`
	SubtopicIDs   []uint               `json:"subtopic_ids" validate:"required"`
}

type CompanyUpdateRequest struct {
	Name          string               `json:"name" validate:"required"`
	Logo          string               `json:"logo"`
	Industry      string               `json:"industry" validate:"required"`
	Website       string               `json:"website"`
	Email         string               `json:"email"`
	Projects      string               `json:"projects"`
	Localiatation *LocalitationRequest `json:"localitation"`
}
