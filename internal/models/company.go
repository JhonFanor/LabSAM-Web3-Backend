package models

import "time"

type Company struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:255;not null"`
	Industry   string `gorm:"size:100"`
	LocationID uint
	Website    string `gorm:"size:255"`
	Email      string `gorm:"size:255"`
	UserID     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Company) TableName() string {
	return "companies"
}
