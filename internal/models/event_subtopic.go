package models

type EventSubtopic struct {
	EventID    uint `gorm:"primaryKey"`
	SubtopicID uint `gorm:"primaryKey"`
}

func (EventSubtopic) TableName() string {
	return "event_subtopic"
}
