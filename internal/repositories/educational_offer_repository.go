package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EducationalOfferRepository interface {
	Create(educationalOffer *models.EducationalOffer) (*models.EducationalOffer, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.EducationalOffer, error)
	Update(educationalOffer *models.EducationalOffer) error
	Delete(id uint) error
}

type educationalOfferRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewEducationalOfferRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) EducationalOfferRepository {
	return &educationalOfferRepository{
		db: db,
		qm: qm,
	}
}

func (r *educationalOfferRepository) Create(educationalOffer *models.EducationalOffer) (*models.EducationalOffer, error) {
	if err := r.db.Create(educationalOffer).Error; err != nil {
		return nil, err
	}
	return educationalOffer, nil
}

func (r *educationalOfferRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.EducationalOffer{})

	return paginationInfo, nil
}

func (r *educationalOfferRepository) GetByID(id uint) (*models.EducationalOffer, error) {
	var educationalOffer models.EducationalOffer
	if err := r.db.First(&educationalOffer, id).Error; err != nil {
		return nil, err
	}
	return &educationalOffer, nil
}

func (r *educationalOfferRepository) Update(educationalOffer *models.EducationalOffer) error {
	return r.db.Save(educationalOffer).Error
}

func (r *educationalOfferRepository) Delete(id uint) error {
	return r.db.Delete(&models.EducationalOffer{}, id).Error
}
