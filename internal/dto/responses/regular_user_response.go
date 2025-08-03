package responses

import (
	"lamsam-web3-backend/internal/models"
	"time"
)

type RegularUserMinimalResponse struct {
	Name string `json:"name"`
}

type RegularUserGetResponse struct {
	Name      string           `json:"name"`
	BirthDate *time.Time       `json:"birth_date"`
	Location  *models.Location `json:"location"`
	Contact   *models.Contact  `json:"contact"`
}
