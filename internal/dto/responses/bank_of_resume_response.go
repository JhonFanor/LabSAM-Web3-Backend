package responses

type BankOfResumeGetAllResponse struct {
	ID    uint                `json:"id"`
	Title string              `json:"title"`
	User  UserMinimalResponse `json:"user"`
}

type BankOfResumeGetResponse struct {
	ID          uint   `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Skills      string `gorm:"not null"`
	Experience  string
	Education   string
	User        UserMinimalResponse       `json:"user,omitempty"`
	Subtopics   []SubtopicMinimalResponse `json:"subtopics,omitempty"`
}
