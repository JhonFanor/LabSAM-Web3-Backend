package requests

type EducationalOfferRequest struct {
	Title       string  `json:"title" validate:"required"`
	Institution string  `json:"institution" validate:"required"`
	Duration    string  `json:"duration" validate:"required"`
	Cost        float64 `json:"cost" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Website     string  `json:"website" validate:"required"`
	SubtopicIDs []uint  `json:"subtopic_ids" validate:"required"`
}
