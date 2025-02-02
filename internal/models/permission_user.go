package models

type PermissionUser struct {
	PermissionID uint `gorm:"primary_key" json:"permission_id"`
	UserID       uint `gorm:"primary_key" json:"user_id"`
}
