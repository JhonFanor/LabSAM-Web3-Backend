package responses

import "lamsam-web3-backend/internal/models"

type UniversityUserMinimalResponse struct {
	Name string `json:"name"`
}

type UniversityUserGetResponse struct {
	Name           string                 `json:"name"`
	UniversityType *models.UniversityType `json:"university_type"`
	Location       *models.Location       `json:"location"`
	Contact        *models.Contact        `json:"contact"`
}
