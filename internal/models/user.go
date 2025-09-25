package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                     uint            `json:"id"`
	Email                  string          `json:"email"`
	Avatar                 string          `json:"avatar"`
	Password               string          `json:"-"`
	EmailVerified          bool            `json:"email_verified"`
	EmailVerificationToken string          `json:"email_verification_token"`
	RoleID                 uint            `json:"-"`
	Role                   Role            `gorm:"foreignkey:RoleID" json:"-"`
	Permissions            []Permission    `gorm:"many2many:permission_user;" json:"-"`
	CreatedAt              *time.Time      `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt              *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	RegularUser            *RegularUser    `gorm:"foreignkey:UserID" json:"regular_user,omitempty"`
	UniversityUser         *UniversityUser `gorm:"foreignkey:UserID" json:"university_user,omitempty"`
	BusinessUser           *BusinessUser   `gorm:"foreignkey:UserID" json:"business_user,omitempty"`
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
