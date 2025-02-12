package models

type DocumentationSubtopic struct {
	DocumentationID uint `gorm:"primaryKey"`
	SubtopicID      uint `gorm:"primaryKey"`
}

func (DocumentationSubtopic) TableName() string {
	return "documentation_subtopic"
}
