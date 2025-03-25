package models

import "time"

type News struct {
	ID          uint       `gorm:"primaryKey" json:"id,omitempty"`
	Title       string     `gorm:"size:255;not null" json:"title,omitempty"`
	Description string     `gorm:"not null" json:"description,omitempty"`
	Image       string     `gorm:"size:255;not null" json:"image,omitempty"`
	Link        string     `json:"link,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	UserID      uint       `json:"user_id,omitempty"`
	User        User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:news_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (News) TableName() string {
	return "news"
}
