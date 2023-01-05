package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Manager struct {
	gorm.Model
	UserName           string `gorm:"type:varchar(12) ;not null"  form:"userName"`
	PassWord           string `gorm:"type:varchar(12);not null" form:"passWord"`
	jwt.StandardClaims `gorm:"-"`
}
