package models

type LegislationSubtopic struct {
	LegislationID uint `gorm:"primaryKey"`
	SubtopicID    uint `gorm:"primaryKey"`
}

func (LegislationSubtopic) TableName() string {
	return "legislation_subtopic"
}
