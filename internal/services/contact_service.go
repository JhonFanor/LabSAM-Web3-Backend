package services

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/repositories"
	"lamsam-web3-backend/internal/utils"
)

type ContactService interface {
	CreateContact(contact *models.Contact) (*models.Contact, error)
	UpdateContact(contact *models.Contact) error
}

type contactService struct {
	repo repositories.ContactRepository
}

func NewContactService(repo repositories.ContactRepository) ContactService {
	return &contactService{
		repo: repo,
	}
}

func (s *contactService) CreateContact(contact *models.Contact) (*models.Contact, error) {
	if contact == nil {
		return nil, customerrors.ErrInvalidData
	}
	return s.repo.Create(contact)
}

func (s *contactService) UpdateContact(contact *models.Contact) error {
	if contact == nil || contact.ID == 0 {
		return customerrors.ErrInvalidData
	}

	updates := utils.StructToMap(contact)
	if len(updates) == 0 {
		return customerrors.ErrNoUpdates
	}

	return s.repo.Update(contact, updates)
}
