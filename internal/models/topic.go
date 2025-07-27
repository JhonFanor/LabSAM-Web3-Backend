package models

import "time"

type Topic struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"size:255" json:"name"`
	Subtopics []Subtopic `gorm:"foreignKey:TopicID" json:"subtopic"`
	CreatedAt *time.Time `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
