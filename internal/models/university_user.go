package models

type UniversityUser struct {
	Name   string `gorm:"size:255"`
	UserID uint   `gorm:"primaryKey"`
}
