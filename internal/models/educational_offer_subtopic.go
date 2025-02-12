package models

type EducationalOfferSubtopic struct {
	EducationalOfferID uint `gorm:"primaryKey"`
	SubtopicID         uint `gorm:"primaryKey"`
}

func (EducationalOfferSubtopic) TableName() string {
	return "educational_offer_subtopic"
}
