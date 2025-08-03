package models

import "time"

type RegularUser struct {
	UserID     uint       `gorm:"primaryKey" json:"user_id"`
	Name       string     `json:"name"`
	BirthDate  *time.Time `json:"birth_date"`
	LocationID *uint      `json:"location_id"`
	ContactID  *uint      `json:"contact_id"`
	Location   *Location  `gorm:"foreignKey:LocationID" json:"location"`
	Contact    *Contact   `gorm:"foreignKey:ContactID" json:"contact"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (RegularUser) TableName() string {
	return "regulars_users"
}
