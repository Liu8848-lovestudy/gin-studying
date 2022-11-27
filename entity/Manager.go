package entity

import "github.com/jinzhu/gorm"

type Manager struct {
	gorm.Model
	UserName string `gorm:"type:varchar(12) ;not null"  form:"userName"`
	PassWord string `gorm:"type:varchar(12);not null" form:"passWord"`
}
