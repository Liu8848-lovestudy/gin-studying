package utils

import "gin-studying/entity"

func CreateTables() {
	db := GetConn()
	db.AutoMigrate(&entity.Book{})
	db.AutoMigrate(&entity.Student{})
	db.AutoMigrate(&entity.BookLent{})
	db.AutoMigrate(&entity.Manager{})
	db.Close()
}
