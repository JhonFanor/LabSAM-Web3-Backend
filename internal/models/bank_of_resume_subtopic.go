package models

type BankOfResumeSubtopic struct {
	BankOfResumeID uint `gorm:"primaryKey"`
	SubtopicID     uint `gorm:"primaryKey"`
}

func (BankOfResumeSubtopic) TableName() string {
	return "bank_of_resume_subtopic"
}
