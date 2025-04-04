package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
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
}

func NewJobBoardRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) JobBoardRepository {
	return &jobBoardRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

func (r *jobBoardRepository) Create(jobBoard *models.JobBoard) (*models.JobBoard, error) {
	if err := r.dbManager.Create(jobBoard); err != nil {
		return nil, err
	}
	return jobBoard, nil
}

func (r *jobBoardRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	r.qm.DB = r.qm.DB.Preload("User")
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.JobBoard{})

	return paginationInfo, nil
}

func (r *jobBoardRepository) GetByID(id uint) (*models.JobBoard, error) {
	r.dbManager.DB = r.qm.DB.Preload("Subtopics").Preload("User")

	var jobBoard models.JobBoard
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&jobBoard, conditions); err != nil {
		return nil, err
	}
	return &jobBoard, nil
}

func (r *jobBoardRepository) Update(jobBoard *models.JobBoard, updates map[string]interface{}) error {
	return r.dbManager.Update(jobBoard, updates)
}

func (r *jobBoardRepository) Delete(id uint) error {
	jobBoard := models.JobBoard{ID: id}
	return r.dbManager.Delete(&jobBoard)
}
