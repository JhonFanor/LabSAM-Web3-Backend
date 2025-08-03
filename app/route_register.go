package app

import (
	"lamsam-web3-backend/internal/api/routes"

	"go.uber.org/fx"
)

type RoutesRegisterParams struct {
	fx.In
	AuthRoutes                  *routes.AuthRoutes
	BankOfResumeRoutes          *routes.BankOfResumeRoutes
	BankOfResumeSubtopicRoutes  *routes.BankOfResumeSubtopicRoutes
	CompanyRoutes               *routes.CompanyRoutes
	CompanySubtopicRoutes       *routes.CompanySubtopicRoutes
	DeniedPermissionUserRoutes  *routes.DeniedPermissionUserRoutes
	DocumentationRoutes         *routes.DocumentationRoutes
	DocumentationSubtopicRoutes *routes.DocumentationSubtopicRoutes
	EducationalOfferRoutes      *routes.EducationalOfferRoutes
	EducationalOfferSubtopic    *routes.EducationalOfferSubtopicRoutes
	EmailVerificationRoutes     *routes.EmailVerificationRoutes
	EventRoutes                 *routes.EventRoutes
	EventSubtopicRoutes         *routes.EventSubtopicRoutes
	InvestigationRoutes         *routes.InvestigationRoutes
	InvestigationSubtopicRoutes *routes.InvestigationSubtopicRoutes
	JobBoardRoutes              *routes.JobBoardRoutes
	JobBoardSubtopicRoutes      *routes.JobBoardSubtopicRoutes
	LegislationRoutes           *routes.LegislationRoutes
	LegislationSubtopicRoutes   *routes.LegislationSubtopicRoutes
	NewsRoutes                  *routes.NewsRoutes
	NewsSubtopicRoutes          *routes.NewsSubtopicRoutes
	NotificationRoutes          *routes.NotificationRoutes
	PermissionRoleRoutes        *routes.PermissionRoleRoutes
	PermissionRoutes            *routes.PermissionRoutes
	PermissionUserRoutes        *routes.PermissionUserRoutes
	Publications                *routes.PublicationRoutes
	RejectionCommentRoutes      *routes.RejectionCommentRoutes
	TopicRoutes                 *routes.TopicRoutes
	UniversityTypeRoutes        *routes.UniversityTypeRoutes
	UploadRoutes                *routes.UploadRoutes
	UserRoutes                  *routes.UserRoutes
}

func RoutesRegister(p RoutesRegisterParams) {
	p.AuthRoutes.Routes()
	p.BankOfResumeRoutes.Routes()
	p.BankOfResumeSubtopicRoutes.Routes()
	p.CompanyRoutes.Routes()
	p.CompanySubtopicRoutes.Routes()
	p.DeniedPermissionUserRoutes.Routes()
	p.DocumentationRoutes.Routes()
	p.DocumentationSubtopicRoutes.Routes()
	p.EducationalOfferRoutes.Routes()
	p.EducationalOfferSubtopic.Routes()
	p.EmailVerificationRoutes.Routes()
	p.EventRoutes.Routes()
	p.EventSubtopicRoutes.Routes()
	p.InvestigationRoutes.Routes()
	p.InvestigationSubtopicRoutes.Routes()
	p.JobBoardRoutes.Routes()
	p.JobBoardSubtopicRoutes.Routes()
	p.LegislationRoutes.Routes()
	p.LegislationSubtopicRoutes.Routes()
	p.NewsRoutes.Routes()
	p.NewsSubtopicRoutes.Routes()
	p.NotificationRoutes.Routes()
	p.PermissionRoleRoutes.Routes()
	p.PermissionRoutes.Routes()
	p.PermissionUserRoutes.Routes()
	p.Publications.Routes()
	p.RejectionCommentRoutes.Routes()
	p.TopicRoutes.Routes()
	p.UniversityTypeRoutes.Routes()
	p.UploadRoutes.Routes()
	p.UserRoutes.Routes()
}
