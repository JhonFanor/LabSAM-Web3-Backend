package repositories

import (
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type EducationalOfferSubtopicRepository interface {
	Create(educationalOfferSubtopic *models.EducationalOfferSubtopic) (*models.EducationalOfferSubtopic, error)
	GetByID(educationalOfferID, subtopicID uint) (*models.EducationalOfferSubtopic, error)
	Delete(educationalOfferID, subtopicID uint) error
}

type educationalOfferSubtopicRepository struct {
	db *gorm.DB
}

func NewEducationalOfferSubtopicRepository(db *gorm.DB) EducationalOfferSubtopicRepository {
	return &educationalOfferSubtopicRepository{
		db: db,
	}
}

func (r *educationalOfferSubtopicRepository) Create(educationalOfferSubtopic *models.EducationalOfferSubtopic) (*models.EducationalOfferSubtopic, error) {
	if err := r.db.Create(educationalOfferSubtopic).Error; err != nil {
		return nil, err
	}
	return educationalOfferSubtopic, nil
}

func (r *educationalOfferSubtopicRepository) GetByID(educationalOfferID, subtopicID uint) (*models.EducationalOfferSubtopic, error) {
	var educationalOfferSubtopic models.EducationalOfferSubtopic
	if err := r.db.Where("educational_offer_id = ? AND subtopic_id = ?", educationalOfferID, subtopicID).First(&educationalOfferSubtopic).Error; err != nil {
		return nil, err
	}
	return &educationalOfferSubtopic, nil
}

func (r *educationalOfferSubtopicRepository) Delete(educationalOfferID, subtopicID uint) error {
	return r.db.Where("educational_offer_id = ? AND subtopic_id = ?", educationalOfferID, subtopicID).
		Delete(&models.EducationalOfferSubtopic{}).Error
}
