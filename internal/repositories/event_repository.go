package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
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
}

func NewEventRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) EventRepository {
	return &eventRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

func (r *eventRepository) Create(event *models.Event) (*models.Event, error) {
	if err := r.dbManager.Create(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	r.qm.DB = r.qm.DB.Preload("User")
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.Event{})

	return paginationInfo, nil
}

func (r *eventRepository) GetByID(id uint) (*models.Event, error) {
	r.dbManager.DB = r.qm.DB.Preload("Subtopics").Preload("User")

	var event models.Event
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&event, conditions); err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) Update(event *models.Event, updates map[string]interface{}) error {
	return r.dbManager.Update(event, updates)
}

func (r *eventRepository) Delete(id uint) error {
	event := models.Event{ID: id}
	return r.dbManager.Delete(&event)
}
