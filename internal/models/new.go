package models

import "time"

type New struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"not null"`
	Image       string `gorm:"size:255;not null"`
	Source      string
	Link        string
	Date        time.Time
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewsSubtopic struct {
	NewsID     uint `gorm:"primaryKey"`
	SubtopicID uint `gorm:"primaryKey"`
}
