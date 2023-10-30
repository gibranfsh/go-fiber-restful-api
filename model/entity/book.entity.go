package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Author      Author `json:"author" gorm:"foreignKey:AuthorID"` // gorm:"foreignKey:AuthorID"
	AuthorID    uint   `json:"author_id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"` // snake_case
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Copies      int    `json:"copies"`    // Jumlah salinan buku yang tersedia di perpustakaan
	Available   int    `json:"available"` // Jumlah buku yang tersedia untuk dipinjam
}
