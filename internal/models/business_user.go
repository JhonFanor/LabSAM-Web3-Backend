package models

import "time"

type BusinessUser struct {
	UserID     uint       `json:"user_id"`
	Name       string     `json:"name"`
	Industry   string     `json:"industry"`
	LocationID *uint      `json:"location_id"`
	ContactID  *uint      `json:"contact_id"`
	Location   *Location  `json:"location"`
	Contact    *Contact   `json:"contact"`
	CreatedAt  *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
