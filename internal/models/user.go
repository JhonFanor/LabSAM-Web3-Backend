package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Username       string `gorm:"unique;not null" json:"username"`
	Email          string `gorm:"unique;not null" json:"email"`
	Password       string `gorm:"not null" json:"-"`
	RoleId         uint
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Role           Role           `gorm:"foreignKey:RoleID"`
	Permissions    []Permission   `gorm:"many2many:permission_user;"`
	RegularUser    RegularUser    `gorm:"foreignKey:UserID"`
	UniversityUser UniversityUser `gorm:"foreignKey:UserID"`
	BusinessUser   BusinessUser   `gorm:"foreignKey:UserID"`
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
