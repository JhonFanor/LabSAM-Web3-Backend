package requests

type CompanyRequest struct {
	Name           string               `json:"name" validate:"required"`
	Industry       string               `json:"industry" validate:"required"`
	LocalitationID int                  `json:"localitation_id" validate:"required"`
	Website        string               `json:"website" validate:"required"`
	Email          string               `json:"email" validate:"required"`
	Localiatation  *LocalitationRequest `json:"localitation"`
	SubtopicIDs    []uint               `json:"subtopic_ids" validate:"required"`
}

type CompanyUpdateRequest struct {
	Name           string               `json:"name" `
	Industry       string               `json:"industry"`
	LocalitationID int                  `json:"localitation_id"`
	Website        string               `json:"website"`
	Email          string               `json:"email"`
	Localiatation  *LocalitationRequest `json:"localitation"`
}
