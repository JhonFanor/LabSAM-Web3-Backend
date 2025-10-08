package models

type TypeEducation struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (TypeEducation) TableName() string {
	return "type_education"
}
