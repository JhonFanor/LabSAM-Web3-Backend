package models

type JobExchangeSubtopic struct {
	JobExchangeID uint `gorm:"primaryKey"`
	SubtopicID    uint `gorm:"primaryKey"`
}

func (JobExchangeSubtopic) TableName() string {
	return "job_exchange_subtopic"
}
