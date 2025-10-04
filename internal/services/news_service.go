package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/observers"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type NewsService interface {
	CreateNews(news *models.News, userID uint, subtopicIDs []uint) (*models.News, error)
	GetAllNews(c *gin.Context) (*dto.PaginationDTO, error)
	GetAllNewsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error)
	GetAllNewsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error)
	CountNewsNotApproved(role string) (int64, error)
	CountNewsBySubtopic() ([]responses.SubtopicCountResponse, error)
	GetNewsByID(id uint, userID uint, role string) (*models.News, error)
	UpdateNews(news *models.News, userID uint, role string) error
	SetNewsApproval(id uint, approved bool, adminId uint, role string) error
	DeleteNews(id uint, userID uint, role string) error
}

type newsService struct {
	repo                      repositories.NewsRepository
	newsSubtopicService       NewsSubtopicService
	subtopicService           SubtopicService
	adminNotificationObserver *observers.AdminNotificationObserver
	userNotificationObserver  *observers.UserNotificationObserver
}

func NewNewsService(repo repositories.NewsRepository, newsSubtopicService NewsSubtopicService, subtopicService SubtopicService, adminNotificationObserver *observers.AdminNotificationObserver, userNotificationObserver *observers.UserNotificationObserver) NewsService {
	return &newsService{
		repo:                      repo,
		newsSubtopicService:       newsSubtopicService,
		subtopicService:           subtopicService,
		adminNotificationObserver: adminNotificationObserver,
		userNotificationObserver:  userNotificationObserver,
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

	s.adminNotificationObserver.Handle(observers.EventObserver{
		Type:         observers.EventObserverType(observers.Created),
		SenderID:     userID,
		Message:      "Ha creado una nueva noticia.",
		Action:       "created",
		ResourceID:   int(createdNews.ID),
		ResourceType: "news",
	})

	return createdNews, nil
}

func (s *newsService) GetAllNews(c *gin.Context) (*dto.PaginationDTO, error) {
	return s.repo.GetAll(c)
}

func (s *newsService) GetAllNewsByUserID(c *gin.Context, userID uint) (*dto.PaginationDTO, error) {
	return s.repo.GetAllByUserID(c, userID)
}

func (s *newsService) GetAllNewsNotApproved(c *gin.Context, role string) (*dto.PaginationDTO, error) {
	if role != "admin" {
		return nil, customerrors.ErrUnauthorized
	}
	return s.repo.GetAllNotApproved(c)
}

func (s *newsService) CountNewsNotApproved(role string) (int64, error) {
	if role != "admin" {
		return 0, customerrors.ErrUnauthorized
	}
	return s.repo.CountNotApproved()
}

func (s *newsService) CountNewsBySubtopic() ([]responses.SubtopicCountResponse, error) {
	subtopics, err := s.subtopicService.GetAllSubtopic()
	if err != nil {
		return nil, err
	}

	var result []responses.SubtopicCountResponse

	for _, sub := range subtopics {
		count, err := s.repo.CountBySubtopicID(sub.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, responses.SubtopicCountResponse{
			SubtopicName: sub.Name,
			Count:        count,
		})
	}

	return result, nil
}

func (s *newsService) GetNewsByID(id uint, userID uint, role string) (*models.News, error) {
	if id == 0 {
		return nil, customerrors.ErrInvalidID
	}

	news, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if news.IsApproved != nil && *news.IsApproved {
		return news, nil
	}

	if news.UserID == userID || role == "admin" {
		return news, nil
	}

	return nil, customerrors.ErrUnauthorized
}

func (s *newsService) UpdateNews(news *models.News, userID uint, role string) error {
	if news == nil || news.ID == 0 {
		return customerrors.ErrInvalidData
	}

	existing, err := s.GetNewsByID(news.ID, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	news.IsApproved = nil

	updates := utils.StructToMap(news)
	if len(updates) == 0 {
		return nil
	}

	if role != "admin" {
		s.adminNotificationObserver.Handle(observers.EventObserver{
			Type:         observers.EventObserverType(observers.Updated),
			SenderID:     userID,
			Message:      "Ha actualizado una noticia.",
			Action:       "updated",
			ResourceID:   int(existing.ID),
			ResourceType: "news",
		})
	}

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *newsService) SetNewsApproval(id uint, approved bool, adminId uint, role string) error {
	if role != "admin" {
		return customerrors.ErrUnauthorized
	}

	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	isApproved := approved
	updates := map[string]interface{}{
		"is_approved": &isApproved,
	}

	message := ""
	action := ""
	var typeObserver observers.EventObserverType
	if approved {
		typeObserver = observers.EventObserverType(observers.Approved)
		message = "El administrador aprobo la noticia."
		action = "approved"
	} else {
		typeObserver = observers.EventObserverType(observers.Rejected)
		message = "El administrador rechazo la noticia."
		action = "rejected"
	}

	s.userNotificationObserver.Handle(observers.EventObserver{
		Type:         typeObserver,
		SenderID:     adminId,
		ReceiverID:   existing.UserID,
		Message:      message,
		Action:       action,
		ResourceID:   int(existing.ID),
		ResourceType: "news",
	})

	existing.User = nil
	return s.repo.Update(existing, updates)
}

func (s *newsService) DeleteNews(id uint, userID uint, role string) error {
	if id == 0 {
		return customerrors.ErrInvalidID
	}

	existing, err := s.GetNewsByID(id, userID, role)
	if err != nil {
		return err
	}

	if existing.UserID != userID && role != "admin" {
		return customerrors.ErrUnauthorized
	}

	return s.repo.Delete(id)
}
