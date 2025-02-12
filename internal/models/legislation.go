package models

import "time"

type Legislation struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"not null"`
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
