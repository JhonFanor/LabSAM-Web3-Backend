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
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	CountBySubtopicID(subtopicID uint) (int64, error)
	GetByID(id uint) (*models.EducationalOffer, error)
	GetRandomApproved() (*models.EducationalOffer, error)
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
	if err := r.dbManager.Create(educationalOffer, r.db); err != nil {
		return nil, err
	}
	return educationalOffer, nil
}

func (r *educationalOfferRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.EducationalOffer{})

	return paginationInfo, nil
}

func (r *educationalOfferRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.EducationalOffer{},
	)

	return paginationInfo, nil
}

func (r *educationalOfferRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.EducationalOffer{},
	)

	return paginationInfo, nil
}

func (r *educationalOfferRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.EducationalOffer{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *educationalOfferRepository) CountBySubtopicID(subtopicID uint) (int64, error) {
	var count int64
	err := r.db.
		Table("educational_offer_subtopic").
		Where("subtopic_id = ?", subtopicID).
		Count(&count).Error
	return count, err
}

func (r *educationalOfferRepository) GetByID(id uint) (*models.EducationalOffer, error) {
	var educationalOffer models.EducationalOffer
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&educationalOffer, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
		return nil, err
	}
	return &educationalOffer, nil
}

func (r *educationalOfferRepository) GetRandomApproved() (*models.EducationalOffer, error) {
	var educationalOffer models.EducationalOffer
	err := r.db.
		Preload("User").
		Where("is_approved = ?", true).
		Order("RANDOM()").
		First(&educationalOffer).Error
	if err != nil {
		return nil, err
	}
	return &educationalOffer, nil
}

func (r *educationalOfferRepository) Update(educationalOffer *models.EducationalOffer, updates map[string]interface{}) error {
	return r.dbManager.Update(educationalOffer, updates, r.db)
}

func (r *educationalOfferRepository) Delete(id uint) error {
	educationalOffer := models.EducationalOffer{ID: id}
	return r.dbManager.Delete(&educationalOffer, r.db)
}
