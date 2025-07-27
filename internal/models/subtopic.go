package models

import "time"

type Subtopic struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	TopicID   uint
	CreatedAt *time.Time `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
