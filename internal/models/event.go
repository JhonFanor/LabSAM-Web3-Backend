package models

import "time"

type Event struct {
	ID             uint          `json:"id"`
	Title          string        `json:"title"`
	Image          string        `json:"image"`
	Description    string        `json:"description"`
	Link           string        `json:"link"`
	Date           time.Time     `json:"date"`
	IsApproved     *bool         `json:"is_approved,omitempty"`
	UserID         uint          `json:"-"`
	LocalitationID *uint         `json:"-"`
	User           *User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Localitation   *Localitation `gorm:"foreignKey:LocalitationID" json:"localitation,omitempty"`
	Subtopics      []Subtopic    `gorm:"many2many:event_subtopic;" json:"subtopics,omitempty"`
	CreatedAt      *time.Time    `json:"created_at,omitempty"`
	UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
}
