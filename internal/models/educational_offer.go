package models

import "time"

type EducationalOffer struct {
	ID              uint          `json:"id"`
	Title           string        `json:"title"`
	Institution     string        `json:"institution"`
	Logo            string        `json:"logo"`
	StartDate       *time.Time    `json:"start_date"`
	EndDate         *time.Time    `json:"end_date"`
	Cost            float64       `json:"cost"`
	Description     string        `json:"description"`
	Link            string        `json:"link"`
	IsApproved      *bool         `json:"is_approved,omitempty"`
	UserID          uint          `json:"-"`
	TypeEducationID uint          `json:"type_education_id"`
	CurrencyTypeID  *uint         `json:"currency_type_id"`
	User            *User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TypeEducation   TypeEducation `gorm:"foreignKey:TypeEducationID" json:"type_education"`
	CurrencyType    CurrencyType  `gorm:"fo0reignKey:CurrencyTypeID" json:"currency_type"`
	Subtopics       []Subtopic    `gorm:"many2many:educational_offer_subtopic;" json:"subtopics,omitempty"`
	CreatedAt       *time.Time    `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	UpdatedAt       *time.Time    `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
