package models

type InvestigationSubtopic struct {
	InvestigationID uint `gorm:"primaryKey"`
	SubtopicID      uint `gorm:"primaryKey"`
}

func (InvestigationSubtopic) TableName() string {
	return "investigation_subtopic"
}
