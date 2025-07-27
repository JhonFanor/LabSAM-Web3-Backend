package models

import (
	"time"
)

type RejectionComment struct {
	ID           uint       `json:"id"`
	ResourceType string     `json:"resource_type"`
	ResourceID   uint       `json:"resource_id"`
	Comment      string     `json:"comment"`
	CreatedAt    *time.Time `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
