package models

type JobBoardSubtopic struct {
	JobBoardID uint `gorm:"primaryKey"`
	SubtopicID uint `gorm:"primaryKey"`
}

func (JobBoardSubtopic) TableName() string {
	return "job_exchange_subtopic"
}
