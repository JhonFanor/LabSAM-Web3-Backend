package models

import "time"

type JobBoard struct {
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

func (JobBoard) TableName() string {
	return "jobs_board"
}
