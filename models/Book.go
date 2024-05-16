package models

import (
	"gorm.io/gorm"
)

type Book struct{
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	YearOfPublication string `json:"year_of_publication"`
}
