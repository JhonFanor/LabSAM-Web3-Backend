package models

import "time"

type Legislation struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	IsApproved  *bool      `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:legislation_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
