package models

type PermissionUser struct {
	PermissionID uint `gorm:"primaryKey"`
	UserID       uint `gorm:"primaryKey"`
}
