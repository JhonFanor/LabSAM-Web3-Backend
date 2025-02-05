package models

type UniversityUser struct {
	Name   string `gorm:"size:255;not null" json:"name"`
	UserID uint   `gorm:"not null;unique" json:"user_id"`
}

func (UniversityUser) TableName() string {
	return "universities_users"
}
