package models

type BusinessUser struct {
	Name   string `gorm:"size:255"`
	UserID uint   `gorm:"primaryKey"`
}
