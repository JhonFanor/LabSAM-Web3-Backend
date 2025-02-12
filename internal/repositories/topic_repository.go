package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/models"

	"gorm.io/gorm"
)

type TopicRepository interface {
	Create(topic *models.Topic) (*models.Topic, error)
	Update(topic *models.Topic) error
	Delete(id uint) error
	GetByID(id uint) (*models.Topic, error)
	GetAll() ([]models.Topic, error)
}

type topicRepository struct {
	db *gorm.DB
	qm *gormmanagers.GormQueryManager
}

func NewTopicRepository(db *gorm.DB, qm *gormmanagers.GormQueryManager) TopicRepository {
	return &topicRepository{
		db: db,
		qm: qm,
	}
}

func (r *topicRepository) Create(topic *models.Topic) (*models.Topic, error) {
	if err := r.db.Create(topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *topicRepository) Update(topic *models.Topic) error {
	return r.db.Save(topic).Error
}

func (r *topicRepository) Delete(id uint) error {
	return r.db.Delete(&models.Topic{}, id).Error
}

func (r *topicRepository) GetByID(id uint) (*models.Topic, error) {
	var topic models.Topic
	if err := r.db.First(&topic, id).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *topicRepository) GetAll() ([]models.Topic, error) {
	var topics []models.Topic

	if err := r.db.Preload("Subtopics").Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}
