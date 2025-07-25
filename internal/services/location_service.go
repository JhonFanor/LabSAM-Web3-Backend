package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
)

type LocationService interface {
	CreateLocation(location *models.Location) (*models.Location, error)
	GetLocationByCountryAndCity(country string, city string) (*models.Location, error)
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{
		repo: repo,
	}
}

func (s *locationService) CreateLocation(location *models.Location) (*models.Location, error) {
	if location == nil {
		return nil, customerrors.ErrInvalidData
	}

	existed, _ := s.GetLocationByCountryAndCity(location.Country, location.City)

	if existed != nil {
		return existed, nil
	}

	return s.repo.Create(location)
}

func (s *locationService) GetLocationByCountryAndCity(country, city string) (*models.Location, error) {
	if country == "" || city == "" {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.GetByCountryAndCity(country, city)
}
