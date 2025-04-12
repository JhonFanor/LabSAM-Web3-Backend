package models

import (
	"time"
)

type Localitation struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Address   string  `gorm:"type:varchar(255);not null"`
	Latitude  float64 `gorm:"type:decimal(9,6)"`
	Longitude float64 `gorm:"type:decimal(9,6)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
