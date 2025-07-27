package models

import (
	"time"
)

type Permission struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Roles     []Role    `gorm:"many2many:permission_rol;" json:"roles"`
	Users     []User    `gorm:"many2many:permission_user;" json:"users"`
}
