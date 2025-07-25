package models

type Contact struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
