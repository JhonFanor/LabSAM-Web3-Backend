package models

type PermissionRole struct {
	PermissionID uint `gorm:"primaryKey"`
	RoleID       uint `gorm:"primaryKey"`
}
