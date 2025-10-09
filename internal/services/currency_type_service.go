package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
)

type CurrencyTypeService interface {
	CreateCurrencyType(currencyType *models.CurrencyType) (*models.CurrencyType, error)
	GetCurrencyTypeByCode(code string) (*models.CurrencyType, error)
}

type currencyTypeService struct {
	repo repositories.CurrencyTypeRepository
}

func NewCurrencyTypeService(repo repositories.CurrencyTypeRepository) CurrencyTypeService {
	return &currencyTypeService{
		repo: repo,
	}
}

func (s *currencyTypeService) CreateCurrencyType(currencyType *models.CurrencyType) (*models.CurrencyType, error) {
	if currencyType == nil {
		return nil, customerrors.ErrInvalidData
	}

	existed, _ := s.GetCurrencyTypeByCode(currencyType.Code)

	if existed != nil {
		return existed, nil
	}

	return s.repo.Create(currencyType)
}

func (s *currencyTypeService) GetCurrencyTypeByCode(code string) (*models.CurrencyType, error) {
	if code == "" {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.GetByCode(code)
}
