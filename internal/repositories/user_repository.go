package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User, updates map[string]interface{}) error
	Delete(id uint) error
	GetByID(id uint) (*models.User, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetAdmins() ([]models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	dbManager *gormmanagers.DBManager
	db        *gorm.DB
	qm        *gormmanagers.GormQueryManager
}

func NewUserRepository(dbManager *gormmanagers.DBManager, db *gorm.DB, qm *gormmanagers.GormQueryManager) UserRepository {
	return &userRepository{
		dbManager: dbManager,
		db:        db,
		qm:        qm,
	}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User, updates map[string]interface{}) error {
	return r.dbManager.Update(user, updates, r.db)
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("RegularUser").
		Preload("UniversityUser").
		Preload("BusinessUser").
		Preload("RegularUser.Location").
		Preload("RegularUser.Contact").
		Preload("UniversityUser.UniversityType").
		Preload("UniversityUser.Location").
		Preload("UniversityUser.Contact").
		Preload("BusinessUser.Location").
		Preload("BusinessUser.Contact").
		First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(
		c,
		r.db.Joins("JOIN roles ON roles.id = users.role_id").
			Where("roles.name <> ?", "admin").
			Preload("RegularUser").
			Preload("UniversityUser").
			Preload("BusinessUser"),
		&models.User{},
	)

	return paginationInfo, nil
}

func (r *userRepository) GetAdmins() ([]models.User, error) {
	var admins []models.User

	err := r.db.
		Joins("JOIN roles ON roles.id = users.role_id").
		Where("roles.name = ?", "admin").
		Preload("Role").
		Find(&admins).Error

	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
