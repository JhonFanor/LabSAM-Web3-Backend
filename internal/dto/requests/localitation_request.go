package requests

type LocalitationRequest struct {
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
