package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) (*models.Event, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.Event, error)
	Update(event *models.Event, updates map[string]interface{}) error
	Delete(id uint) error
}

type eventRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
	db        *gorm.DB
}

func NewEventRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager, db *gorm.DB) EventRepository {
	return &eventRepository{
		dbManager: dbManager,
		qm:        qm,
		db:        db,
	}
}

func (r *eventRepository) Create(event *models.Event) (*models.Event, error) {
	if err := r.dbManager.Create(event, r.db); err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, r.db.Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Where("is_approved = ?", true), &models.Event{})

	return paginationInfo, nil
}

func (r *eventRepository) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&event, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Preload("Localitation")); err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) Update(event *models.Event, updates map[string]interface{}) error {
	return r.dbManager.Update(event, updates, r.db)
}

func (r *eventRepository) Delete(id uint) error {
	event := models.Event{ID: id}
	return r.dbManager.Delete(&event, r.db)
}
