package services

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type EducationalOfferService interface {
	CreateEducationalOffer(educationalOffer *models.EducationalOffer, userID uint, subtopicIDs []uint) (*models.EducationalOffer, error)
}

type educationalOfferService struct {
	repo                            repositories.EducationalOfferRepository
	educationalOfferSubtopicService EducationalOfferSubtopicService
	db                              *gorm.DB
	qm                              *gormmanagers.GormQueryManager
}

func NewEducationalOfferService(repo repositories.EducationalOfferRepository, educationalOfferSubtopicService EducationalOfferSubtopicService, db *gorm.DB, qm *gormmanagers.GormQueryManager) EducationalOfferService {
	return &educationalOfferService{
		repo:                            repo,
		educationalOfferSubtopicService: educationalOfferSubtopicService,
		db:                              db,
		qm:                              qm,
	}
}

func (s *educationalOfferService) CreateEducationalOffer(educationalOffer *models.EducationalOffer, userID uint, subtopicIDs []uint) (*models.EducationalOffer, error) {
	if educationalOffer == nil {
		return nil, customerrors.ErrInvalidData
	}

	educationalOffer.UserID = userID

	createdEducationalOffer, err := s.repo.Create(educationalOffer)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		educationalOfferSubtopic := &models.EducationalOfferSubtopic{
			EducationalOfferID: createdEducationalOffer.ID,
			SubtopicID:         uint(subtopicID),
		}

		_, err := s.educationalOfferSubtopicService.CreateEducationalOfferSubtopic(educationalOfferSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdEducationalOffer, nil
}
