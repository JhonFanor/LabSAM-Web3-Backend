package models

import (
	"time"
)

type Permission struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Roles     []Role    `gorm:"many2many:permission_rol;" json:"roles"`
	Users     []User    `gorm:"many2many:permission_user;" json:"users"`
}
