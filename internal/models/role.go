package models

import (
	"time"
)

type Role struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `gorm:"many2many:permission_rol;" json:"permissions"`
	Users       []User       `gorm:"foreignkey:RoleID" json:"users"`
	CreatedAt   *time.Time   `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt   *time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
