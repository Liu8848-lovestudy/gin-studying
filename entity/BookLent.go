package entity

import "github.com/jinzhu/gorm"

type BookLent struct {
	StuId  int
	BookId int
	gorm.Model
}
