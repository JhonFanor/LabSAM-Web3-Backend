package repositories

import (
	gormmanagers "lamsam-web3-backend/internal/adapter/gorm/managers"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"

	"github.com/gin-gonic/gin"
)

type NewsRepository interface {
	Create(news *models.News) (*models.News, error)
	GetAll(c *gin.Context) (*dto.PaginationDTO, error)
	GetByID(id uint) (*models.News, error)
	Update(news *models.News, updates map[string]interface{}) error
	Delete(id uint) error
}

type newsRepository struct {
	dbManager *gormmanagers.DBManager
	qm        *gormmanagers.GormQueryManager
}

// Constructor modificado para recibir DBManager
func NewNewsRepository(dbManager *gormmanagers.DBManager, qm *gormmanagers.GormQueryManager) NewsRepository {
	return &newsRepository{
		dbManager: dbManager,
		qm:        qm,
	}
}

// Método para crear una noticia
func (r *newsRepository) Create(news *models.News) (*models.News, error) {
	if err := r.dbManager.Create(news); err != nil {
		return nil, err
	}
	return news, nil
}

// Método para obtener todas las noticias con paginación
func (r *newsRepository) GetAll(c *gin.Context) (*dto.PaginationDTO, error) {
	paginationInfo := r.qm.ApplyPaginationAndFilters(c, &models.News{})
	return paginationInfo, nil
}

// Método para obtener una noticia por ID
func (r *newsRepository) GetByID(id uint) (*models.News, error) {
	var news models.News
	conditions := map[string]interface{}{"id": id}
	if err := r.dbManager.Find(&news, conditions); err != nil {
		return nil, err
	}
	return &news, nil
}

// Método para actualizar una noticia
func (r *newsRepository) Update(news *models.News, updates map[string]interface{}) error {
	return r.dbManager.Update(news, updates)
}

// Método para eliminar una noticia
func (r *newsRepository) Delete(id uint) error {
	news := models.News{ID: id} // Crear instancia solo con el ID
	return r.dbManager.Delete(&news)
}
