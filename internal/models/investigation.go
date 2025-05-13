package models

import "time"

type Investigation struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        time.Time  `json:"date"`
	Link        string     `json:"link"`
	IsApproved  bool       `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:news_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
