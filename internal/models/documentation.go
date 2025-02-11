package models

import "time"

type Documentation struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"not null"`
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DocumentationSubtopic struct {
	DocumentationID uint `gorm:"primaryKey"`
	SubtopicID      uint `gorm:"primaryKey"`
}
