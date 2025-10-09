package models

import "time"

type JobBoard struct {
	ID             uint         `json:"id"`
	Title          string       `json:"title"`
	Logo           string       `json:"logo"`
	Company        string       `json:"company"`
	Description    string       `json:"description"`
	Type           string       `json:"type"`
	SalaryRange    string       `json:"salary_range"`
	Link           string       `json:"link"`
	StartDate      *time.Time   `json:"start_date"`
	EndDate        *time.Time   `json:"end_date"`
	IsApproved     *bool        `json:"is_approved,omitempty"`
	UserID         uint         `json:"-"`
	CurrencyTypeID *uint        `json:"-"`
	User           *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CurrencyType   CurrencyType `gorm:"foreignKey:CurrencyTypeID" json:"currency_type_id"`
	Subtopics      []Subtopic   `gorm:"many2many:job_board_subtopic;" json:"subtopics,omitempty"`
	CreatedAt      *time.Time   `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt      *time.Time   `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

func (JobBoard) TableName() string {
	return "jobs_board"
}
