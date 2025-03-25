package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type NewsService interface {
	CreateNews(news *models.News, userID uint, subtopicIDs []uint) (*models.News, error)
	GetAllNews(c *gin.Context) (*dto.PaginationDTO, error)
	GetNewsByID(id uint) (*models.News, error)
	UpdateNews(news *models.News, userID uint, role string) error
	DeleteNews(id uint, userID uint, role string) error
}

type newsService struct {
	repo                repositories.NewsRepository
	newsSubtopicService NewsSubtopicService
}

func NewNewsService(repo repositories.NewsRepository, newsSubtopicService NewsSubtopicService) NewsService {
	return &newsService{
		repo:                repo,
		newsSubtopicService: newsSubtopicService,
	}
}

func (s *newsService) CreateNews(news *models.News, userID uint, subtopicIDs []uint) (*models.News, error) {
	if news == nil {
		return nil, customerrors.ErrInvalidData
	}

	news.UserID = userID

	createdNews, err := s.repo.Create(news)
	if err != nil {
		return nil, err
	}

	for _, subtopicID := range subtopicIDs {
		newsSubtopic := &models.NewsSubtopic{
			NewsID:     createdNews.ID,
			SubtopicID: subtopicID,
		}

		_, err := s.newsSubtopicService.CreateNewsSubtopic(newsSubtopic)
		if err != nil {
			return nil, err
		}
	}

	return createdNews, nil
}

func (s *newsService) GetAllNews(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *newsService) GetNewsByID(id uint) (*models.News, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

func (s *newsService) UpdateNews(news *models.News, userID uint, role string) error {
	if news == nil || news.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetNewsByID(news.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	updates := utils.StructToMap(news)
	if len(updates) == 0 {
		return customerrors.ErrNoUpdates
	}

	return s.repo.Update(existing, updates)
}

func (s *newsService) DeleteNews(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetNewsByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
