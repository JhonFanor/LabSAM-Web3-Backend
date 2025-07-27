package models

import "time"

type BankOfResume struct {
	ID         uint       `json:"id"`
	Photo      string     `json:"photo"`
	Title      string     `json:"title"`
	Summary    string     `json:"summary"`
	Link       string     `json:"link"`
	IsApproved *bool      `json:"is_approved"`
	UserID     uint       `json:"-"`
	User       *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics  []Subtopic `gorm:"many2many:bank_of_resume_subtopic;" json:"subtopics,omitempty"`
	CreatedAt  *time.Time `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
