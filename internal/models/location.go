package models

type Location struct {
	ID      uint   `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
}
