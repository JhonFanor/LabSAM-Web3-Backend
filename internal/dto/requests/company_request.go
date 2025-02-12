package requests

type CompanyRequest struct {
	Name           string `json:"name" validate:"required"`
	Industry       string `json:"industry" validate:"required"`
	LocalitationID int    `json:"localitation_id" validate:"required"`
	Website        string `json:"website" validate:"required"`
	Email          string `json:"email" validate:"required"`
	SubtopicIDs    []uint `json:"subtopic_ids" validate:"required"`
}
