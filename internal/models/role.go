package models

import (
	"time"
)

type Role struct {
	ID          uint         `gorm:"primary_key;auto_increment" json:"id"`
	Name        string       `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Permissions []Permission `gorm:"many2many:permission_rol;" json:"permissions"`
	Users       []User       `gorm:"foreignkey:RoleID" json:"users"`
}
