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
	GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error)
	CountNotApproved() (int64, error)
	CountBySubtopicID(subtopicID uint) (int64, error)
	GetByID(id uint) (*models.Event, error)
	GetRandomApproved() (*models.Event, error)
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

func (r *eventRepository) GetAllByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("user_id = ?", userID),
		&models.Event{},
	), nil
}

func (r *eventRepository) GetAllNotApproved(c *gin.Context) (*dto.PaginationDTO, error) {
	return r.qm.ApplyPaginationAndFilters(
		c,
		r.db.
			Preload("User").
			Preload("User.RegularUser").
			Preload("User.UniversityUser").
			Preload("User.BusinessUser").
			Where("is_approved IS NULL"),
		&models.Event{},
	), nil
}

func (r *eventRepository) CountNotApproved() (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Event{}).
		Where("is_approved IS NULL").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *eventRepository) CountBySubtopicID(subtopicID uint) (int64, error) {
	var count int64
	err := r.db.
		Table("event_subtopic").
		Where("subtopic_id = ?", subtopicID).
		Count(&count).Error
	return count, err
}

func (r *eventRepository) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&event, conditions, r.db.Preload("Subtopics").Preload("User").Preload("User.RegularUser").Preload("User.UniversityUser").Preload("User.BusinessUser").Preload("Localitation")); err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) GetRandomApproved() (*models.Event, error) {
	var event models.Event
	err := r.db.
		Preload("User").
		Preload("User.RegularUser").
		Preload("User.UniversityUser").
		Preload("User.BusinessUser").
		Where("is_approved = ?", true).
		Order("RANDOM()").
		First(&event).Error
	if err != nil {
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
