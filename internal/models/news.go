package models

import "time"

type News struct {
	ID          uint       `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Image       string     `json:"image,omitempty"`
	Description string     `json:"description,omitempty"`
	Link        string     `json:"link,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	IsApproved  *bool      `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:news_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

func (News) TableName() string {
	return "news"
}
