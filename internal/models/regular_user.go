package models

type RegularUser struct {
	Name   string `gorm:"size:255"`
	UserID uint   `gorm:"primaryKey"`
}
