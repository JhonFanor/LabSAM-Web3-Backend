package models

type UniversityUser struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

func (UniversityUser) TableName() string {
	return "universities_users"
}
