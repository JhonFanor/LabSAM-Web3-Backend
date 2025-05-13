package models

import "time"

type EducationalOffer struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Institution string     `json:"institution"`
	StartDate   string     `json:"start_date"`
	EndDate     string     `json:"end_date"`
	Cost        float64    `json:"cost"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	IsApproved  bool       `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:news_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
