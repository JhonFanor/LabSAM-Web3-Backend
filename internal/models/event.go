package models

import "time"

type Event struct {
	ID             uint   `gorm:"primaryKey"`
	Title          string `gorm:"size:255;not null"`
	Description    string `gorm:"not null"`
	Link           string
	Date           time.Time
	UserID         uint `gorm:"not null"`
	LocalitationID uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
