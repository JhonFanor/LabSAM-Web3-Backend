package models

import (
	"time"
)

type Permission struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Roles     []Role `gorm:"many2many:permission_rol;"`
	Users     []User `gorm:"many2many:permission_user;"`
}
