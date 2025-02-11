package models

import "time"

type JobExchange struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Company     string `gorm:"size:255"`
	Description string `gorm:"not null"`
	Type        string `gorm:"size:50"`
	SalaryRange string `gorm:"size:100"`
	Status      string `gorm:"size:50"`
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type JobExchangeSubtopic struct {
	JobExchangeID uint `gorm:"primaryKey"`
	SubtopicID    uint `gorm:"primaryKey"`
}
