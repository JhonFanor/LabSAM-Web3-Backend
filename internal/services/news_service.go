package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"

	"gorm.io/gorm"
)

type NewsService interface {
	CreateNews(news *models.News, userID uint, subtopicIDs []uint) (*models.News, error)
}

type newsService struct {
	repo                repositories.NewsRepository
	newsSubtopicService NewsSubtopicService
	db                  *gorm.DB
}

func NewNewsService(repo repositories.NewsRepository, newsSubtopicService NewsSubtopicService, db *gorm.DB) NewsService {
	return &newsService{
		repo:                repo,
		newsSubtopicService: newsSubtopicService,
		db:                  db,
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
