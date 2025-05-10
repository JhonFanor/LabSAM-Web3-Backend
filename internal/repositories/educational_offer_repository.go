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
	Update(educationalOffer *models.EducationalOffer, updates map[string]interface{}) error
	Delete(id uint) error
}

type educationalOfferRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewEducationalOfferRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) EducationalOfferRepository {
	return &educationalOfferRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *educationalOfferRepository) Create(educationalOffer *models.EducationalOffer) (*models.EducationalOffer, error) {
	if err := r.dbManager.Create(educationalOffer); err != nil {
		return nil, err
	}
	return educationalOffer, nil
}

func (r *educationalOfferRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Where("is_approved = ?", true), &models.EducationalOffer{})

	return paginationInfo, nil
}

func (r *educationalOfferRepository) GetByID(id uint) (*models.EducationalOffer, error) {
	r.dbManager.DB = r.dbManager.DB.Preload("Subtopics").Preload("User")

	var educationalOffer models.EducationalOffer
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&educationalOffer, conditions); err != nil {
		return nil, err
	}
	return &educationalOffer, nil
}

func (r *educationalOfferRepository) Update(educationalOffer *models.EducationalOffer, updates map[string]interface{}) error {
	return r.dbManager.Update(educationalOffer, updates)
}

func (r *educationalOfferRepository) Delete(id uint) error {
	educationalOffer := models.EducationalOffer{ID: id}
	return r.dbManager.Delete(&educationalOffer)
}
