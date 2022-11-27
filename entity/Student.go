package entity

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	StuName     string `gorm:"type:varchar(50);not null" form:"stuName"`
	StuSex      string `gorm:"type:varchar(1)" form:"stuSex"`
	PhoneNumber string `gorm:"type:varchar(11)" form:"phoneNumber"`
	StuClass    string `gorm:"type:varchar(20)" form:"stuClass"`
}
