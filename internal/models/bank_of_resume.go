package models

import "time"

type BankOfResume struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"not null"`
	Skills      string `gorm:"not null"`
	Experience  string
	Education   string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
