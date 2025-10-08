package models

import "time"

type Investigation struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Description string     `json:"description"`
	Logo        string     `json:"logo"`
	Date        *time.Time `json:"date"`
	Link        string     `json:"link"`
	IsApproved  *bool      `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:investigation_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
