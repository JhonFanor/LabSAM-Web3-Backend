package responses

import "time"

type EducationalOfferGetAllResponse struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Institution string              `json:"institution"`
	Logo        string              `json:"logo"`
	StartDate   *time.Time          `json:"start_date"`
	EndDate     *time.Time          `json:"end_date"`
	Cost        float64             `json:"cost"`
	User        UserMinimalResponse `json:"user"`
}

type EducationalOfferGetAllByUserIDResponse struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Institution string              `json:"institution"`
	Logo        string              `json:"logo"`
	StartDate   *time.Time          `json:"start_date"`
	EndDate     *time.Time          `json:"end_date"`
	Cost        float64             `json:"cost"`
	IsApproved  *bool               `json:"is_approved"`
	User        UserMinimalResponse `json:"user"`
}

type EducationalOfferGetResponse struct {
	ID            uint                  `json:"id"`
	Title         string                `json:"title"`
	Institution   string                `json:"institution"`
	Logo          string                `json:"logo"`
	StartDate     *time.Time            `json:"start_date"`
	EndDate       *time.Time            `json:"end_date"`
	Cost          float64               `json:"cost"`
	Description   string                `json:"description"`
	Link          string                `json:"link"`
	IsApproved    *bool                 `json:"is_approved"`
	User          UserMinimalResponse   `json:"user"`
	TypeEducation TypeEducationResponse `json:"type_education"`
	CurrencyType  CurrencyTypeRespone   `json:"currency_type"`
	Subtopics     []SubtopicGetResponse `json:"subtopics"`
}
