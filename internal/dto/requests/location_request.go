package requests

type LocationRequest struct {
	Country string `json:"country"`
	City    string `json:"city"`
}
