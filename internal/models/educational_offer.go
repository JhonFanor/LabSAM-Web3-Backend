package models

import "time"

type EducationalOffer struct {
	ID          uint    `gorm:"primaryKey"`
	Title       string  `gorm:"size:255;not null"`
	Institution string  `gorm:"size:255;not null"`
	Duration    string  `gorm:"size:100"`
	Cost        float64 `gorm:"not null"`
	Description string  `gorm:"not null"`
	Website     string  `gorm:"size:500"`
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EducationalOfferSubtopic struct {
	EducationalOfferID uint `gorm:"primaryKey"`
	SubtopicID         uint `gorm:"primaryKey"`
}
