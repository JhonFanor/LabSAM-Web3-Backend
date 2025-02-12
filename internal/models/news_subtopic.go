package models

type NewsSubtopic struct {
	NewsID     uint `gorm:"primaryKey"`
	SubtopicID uint `gorm:"primaryKey"`
}

func (NewsSubtopic) TableName() string {
	return "news_subtopic"
}
