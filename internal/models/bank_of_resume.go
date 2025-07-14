package models

import "time"

type BankOfResume struct {
	ID         uint       `json:"id"`
	Photo      string     `json:"photo"`
	Title      string     `json:"title"`
	Summary    string     `json:"summary"`
	Link       string     `json:"link"`
	IsApproved *bool      `json:"is_approved,omitempty"`
	UserID     uint       `json:"-"`
	User       *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics  []Subtopic `gorm:"many2many:bank_of_resume_subtopic;" json:"subtopics,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
