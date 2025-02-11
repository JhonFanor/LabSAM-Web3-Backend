package models

import "time"

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"not null"`
	Link        string
	Date        time.Time
	UserID      uint `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EventSubtopic struct {
	EventID    uint `gorm:"primaryKey"`
	SubtopicID uint `gorm:"primaryKey"`
}
