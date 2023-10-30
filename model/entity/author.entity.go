package entity

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Books     []Book `json:"books"` // `gorm:"foreignKey:AuthorID"``
}
