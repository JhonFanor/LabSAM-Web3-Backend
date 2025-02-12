package models

import "time"

type Topic struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"size:255"`
	Subtopics []Subtopic `gorm:"foreignKey:TopicID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
