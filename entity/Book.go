package entity

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	BoName   string  `gorm:"type:varchar(50);not null" form:"bookName"`
	BoPrice  float64 `gorm:"not null" form:"bookPrice"`
	BoAuthor string  `gorm:"not null" form:"bookAuthor"`
	BoStock  int     `gorm:"not null" form:"bookStock"`
}
