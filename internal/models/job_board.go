package models

import "time"

type JobBoard struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Company     string     `json:"company"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	SalaryRange string     `json:"salary_range"`
	Link        string     `json:"link"`
	IsApproved  *bool      `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Subtopics   []Subtopic `gorm:"many2many:job_board_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (JobBoard) TableName() string {
	return "jobs_board"
}
