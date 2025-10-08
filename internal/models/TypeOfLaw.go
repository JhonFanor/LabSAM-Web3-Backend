package models

type TypeOfLaw struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (TypeOfLaw) TableName() string {
	return "type_of_law"
}
