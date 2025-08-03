package requests

type ContactRequest struct {
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
