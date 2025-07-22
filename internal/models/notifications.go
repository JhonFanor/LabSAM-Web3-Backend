package models

import "time"

type Notification struct {
	ID           uint       `json:"id"`
	SenderID     *uint      `json:"sender_id"`
	ReceiverID   uint       `json:"receiver_id"`
	Message      string     `json:"message"`
	ResourceType *string    `json:"resource_type"`
	ResourceID   *int       `json:"resource_id"`
	Action       string     `json:"action"`
	IsRead       bool       `json:"is_read"`
	Sender       *User      `gorm:"foreignKey:SenderID;" json:"sender"`
	Receiver     User       `gorm:"foreignKey:ReceiverID;" json:"receiver"`
	CreatedAt    *time.Time `json:"created_at"`
}
