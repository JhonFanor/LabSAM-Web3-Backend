package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobBoardRepository interface {
	Create(jobBoard *models.JobBoard) (*models.JobBoard, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.JobBoard, error)
	Update(jobBoard *models.JobBoard, updates map[string]interface{}) error
	Delete(id uint) error
}

type jobBoardRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewJobBoardRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) JobBoardRepository {
	return &jobBoardRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *jobBoardRepository) Create(jobBoard *models.JobBoard) (*models.JobBoard, error) {
	if err := r.dbManager.Create(jobBoard, r.db); err != nil {
		return nil, err
	}
	return jobBoard, nil
}

func (r *jobBoardRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.JobBoard{})

	return paginationInfo, nil
}

func (r *jobBoardRepository) GetByID(id uint) (*models.JobBoard, error) {
	var jobBoard models.JobBoard
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&jobBoard, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser")); err != nil {
		return nil, err
	}
	return &jobBoard, nil
}

func (r *jobBoardRepository) Update(jobBoard *models.JobBoard, updates map[string]interface{}) error {
	return r.dbManager.Update(jobBoard, updates, r.db)
}

func (r *jobBoardRepository) Delete(id uint) error {
	jobBoard := models.JobBoard{ID: id}
	return r.dbManager.Delete(&jobBoard, r.db)
}
