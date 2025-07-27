package models

type Contact struct {
	ID      uint   `json:"id"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
