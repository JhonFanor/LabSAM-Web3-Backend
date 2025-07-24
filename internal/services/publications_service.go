package services

import (
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type PublicationsService interface {
	GetAllPublications(c *gin.Context) ([]responses.PublicationsResponse, error)
}

type publicationsService struct {
	newsRepository             repositories.NewsRepository
	eventRepository            repositories.EventRepository
	investigationRepository    repositories.InvestigationRepository
	jobBoardRepository         repositories.JobBoardRepository
	bankOfResumeRepository     repositories.BankOfResumeRepository
	companyRepository          repositories.CompanyRepository
	educationalOfferRepository repositories.EducationalOfferRepository
	legislationRepository      repositories.LegislationRepository
	documentationRepository    repositories.DocumentationRepository
}

func NewPublicationService(newsRepository repositories.NewsRepository, eventRepository repositories.EventRepository, investigationRepository repositories.InvestigationRepository, jobBoardRepository repositories.JobBoardRepository, bankOfResumeRepository repositories.BankOfResumeRepository, companyRepository repositories.CompanyRepository, educationalOfferRepository repositories.EducationalOfferRepository, legislationRepository repositories.LegislationRepository, documentationRepository repositories.DocumentationRepository) PublicationsService {
	return &publicationsService{
		newsRepository:             newsRepository,
		eventRepository:            eventRepository,
		investigationRepository:    investigationRepository,
		jobBoardRepository:         jobBoardRepository,
		bankOfResumeRepository:     bankOfResumeRepository,
		companyRepository:          companyRepository,
		educationalOfferRepository: educationalOfferRepository,
		legislationRepository:      legislationRepository,
		documentationRepository:    documentationRepository,
	}
}

func (s *publicationsService) GetAllPublications(c *gin.Context) ([]responses.PublicationsResponse, error) {
	var result []responses.PublicationsResponse

	news, err := s.newsRepository.GetRandomApproved()
	if err == nil && news != nil {
		var newsResponse responses.NewsGetAllResponse
		if err := mapstructure.Decode(news, &newsResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "news",
			Data:         newsResponse,
		})
	}

	event, err := s.eventRepository.GetRandomApproved()
	if err == nil && event != nil {
		var eventResponse responses.EventGetAllResponse
		if err := mapstructure.Decode(event, &eventResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "event",
			Data:         eventResponse,
		})
	}

	inv, err := s.investigationRepository.GetRandomApproved()
	if err == nil && inv != nil {
		var invResponse responses.InvestigationGetAllResponse
		if err := mapstructure.Decode(inv, &invResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "investigation",
			Data:         invResponse,
		})
	}

	job, err := s.jobBoardRepository.GetRandomApproved()
	if err == nil && job != nil {
		var jobResponse responses.JobBoardGetAllResponse
		if err := mapstructure.Decode(job, &jobResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "job_board",
			Data:         jobResponse,
		})
	}

	resume, err := s.bankOfResumeRepository.GetRandomApproved()
	if err == nil && resume != nil {
		var resumeResponse responses.BankOfResumeGetAllResponse
		if err := mapstructure.Decode(resume, &resumeResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "bank_of_resume",
			Data:         resumeResponse,
		})
	}

	company, err := s.companyRepository.GetRandomApproved()
	if err == nil && company != nil {
		var companyResponse responses.CompanyGetAllResponse
		if err := mapstructure.Decode(company, &companyResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "company",
			Data:         companyResponse,
		})
	}

	edu, err := s.educationalOfferRepository.GetRandomApproved()
	if err == nil && edu != nil {
		var eduResponse responses.EducationalOfferGetAllResponse
		if err := mapstructure.Decode(edu, &eduResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "educational_offer",
			Data:         eduResponse,
		})
	}

	leg, err := s.legislationRepository.GetRandomApproved()
	if err == nil && leg != nil {
		var legResponse responses.LegislationGetAllResponse
		if err := mapstructure.Decode(leg, &legResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "legislation",
			Data:         legResponse,
		})
	}

	doc, err := s.documentationRepository.GetRandomApproved()
	if err == nil && doc != nil {
		var docResponse responses.DocumentationGetAllResponse
		if err := mapstructure.Decode(doc, &docResponse); err != nil {
			return nil, err
		}
		result = append(result, responses.PublicationsResponse{
			ResourceType: "documentation",
			Data:         docResponse,
		})
	}

	return result, nil
}
