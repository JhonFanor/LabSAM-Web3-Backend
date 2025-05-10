package models

type RegularUser struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

func (RegularUser) TableName() string {
	return "regulars_users"
}
