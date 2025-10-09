package models

import "time"

type Legislation struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Logo        string     `json:"logo"`
	Date        *time.Time `json:"date"`
	Link        string     `json:"link"`
	IsApproved  *bool      `json:"is_approved,omitempty"`
	UserID      uint       `json:"-"`
	TypeOfLawID uint       `json:"type_of_law_id"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TypeOfLaw   *TypeOfLaw `gorm:"foreignKey:TypeOfLawID" json:"type_of_law"`
	Subtopics   []Subtopic `gorm:"many2many:legislation_subtopic;" json:"subtopics,omitempty"`
	CreatedAt   *time.Time `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
