package models

import "time"

type Company struct {
	ID             uint          `json:"id"`
	Name           string        `json:"name"`
	Industry       string        `json:"industry"`
	Website        string        `json:"website"`
	Email          string        `json:"email"`
	IsApproved     bool          `json:"is_approved,omitempty"`
	UserID         uint          `json:"-"`
	LocalitationID uint          `json:"-"`
	User           *User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Localitation   *Localitation `gorm:"foreignKey:LocalitationID" json:"localitation,omitempty"`
	Subtopics      []Subtopic    `gorm:"many2many:news_subtopic;" json:"subtopics,omitempty"`
	CreatedAt      *time.Time    `json:"created_at,omitempty"`
	UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
}

func (Company) TableName() string {
	return "companies"
}
