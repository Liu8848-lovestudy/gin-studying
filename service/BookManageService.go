package service

import (
	"gin-studying/dto"
	"gin-studying/entity"
	"gin-studying/utils"
)

func AddBookService(book entity.Book) dto.Result {
	db := utils.GetConn()
	db.Create(&book)
	db.Close()
	return dto.NewResult(200, "图书添加成功", book)
}

func DeleteBookService(id1 int) dto.Result {
	var book entity.Book
	book.ID = uint(id1)
	db := utils.GetConn()
	db.Debug().Delete(&book)
	db.Close()
	return dto.NewResult(200, "图书删除成功", nil)
}
func ModifyBookService(book entity.Book) dto.Result {
	db := utils.GetConn()
	db.Model(&entity.Book{}).Update(&book)
	db.Close()
	return dto.NewResult(200, "图书信息修改成功", nil)
}
