package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint           `gorm:"primary_key;auto_increment" json:"id"`
	Username       string         `gorm:"size:255;not null;unique" json:"username"`
	Email          string         `gorm:"size:255;not null;unique" json:"email"`
	Avatar         string         `gorm:"size:255;not null" json:"avatar"`
	Password       string         `gorm:"size:255;not null" json:"-"`
	RoleID         uint           `gorm:"not null" json:"-"`
	Role           Role           `gorm:"foreignkey:RoleID" json:"-"`
	Permissions    []Permission   `gorm:"many2many:permission_user;" json:"-"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	RegularUser    RegularUser    `gorm:"foreignkey:UserID" json:"-"`
	UniversityUser UniversityUser `gorm:"foreignkey:UserID" json:"-"`
	BusinessUser   BusinessUser   `gorm:"foreignkey:UserID" json:"-"`
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
