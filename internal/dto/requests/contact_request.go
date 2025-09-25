package requests

type ContactRequest struct {
	ID 		*uint  `json:"id"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
