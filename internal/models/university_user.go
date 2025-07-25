package models

import "time"

type UniversityUser struct {
	UserID           uint            `json:"user_id"`
	Name             string          `json:"name"`
	UniversityTypeID *uint           `json:"university_type_id"`
	LocationID       *uint           `json:"location_id"`
	ContactID        *uint           `json:"contact_id"`
	UniversityType   *UniversityType `json:"university_type"`
	Location         *Location       `json:"location"`
	Contact          *Contact        `json:"contact"`
	CreatedAt        *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UniversityUser) TableName() string {
	return "universities_users"
}
