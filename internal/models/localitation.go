package models

import (
	"time"
)

type Localitation struct {
	ID        uint    `json:"id"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
