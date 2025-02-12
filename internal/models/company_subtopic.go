package models

type CompanySubtopic struct {
	CompanyID  uint `gorm:"primaryKey"`
	SubtopicID uint `gorm:"primaryKey"`
}

func (CompanySubtopic) TableName() string {
	return "company_subtopic"
}
