package responses

import "lamsam-web3-backend/internal/models"

type BusinessUserMinimalResponse struct {
	Name string `json:"name"`
}

type BusinessUserGetResponse struct {
	Name     string           `json:"name"`
	Industry string           `json:"industry"`
	Location *models.Location `json:"location"`
	Contact  *models.Contact  `json:"contact"`
}
