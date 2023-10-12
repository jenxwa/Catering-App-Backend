package models

type Meal struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Meal     string `json:"title"`
	Category string `json:"author"`
}
