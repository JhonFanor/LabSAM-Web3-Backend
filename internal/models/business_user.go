package models

import "time"

type BusinessUser struct {
	UserID     uint       `gorm:"primaryKey" json:"user_id"`
	Name       string     `json:"name"`
	Industry   string     `json:"industry"`
	LocationID *uint      `json:"location_id"`
	ContactID  *uint      `json:"contact_id"`
	Location   *Location  `gorm:"foreignKey:LocationID" json:"location"`
	Contact    *Contact   `gorm:"foreignKey:ContactID" json:"contact"`
	CreatedAt  *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
