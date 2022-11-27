package service

import (
	"gin-studying/dto"
	"gin-studying/entity"
	"gin-studying/utils"
	"github.com/jinzhu/gorm"
)

func LendBookService(id1, id2 int) dto.Result {
	var bl entity.BookLent
	bl.StuId = id1
	bl.BookId = id2
	db := utils.GetConn()
	var book entity.Book
	book.ID = uint(id2)
	db.Find(&book)
	if book.BoStock <= 0 {
		return dto.NewResult(400, "借书失败，库存不足", nil)
	}
	db.Model(&book).Update("bo_stock", gorm.Expr("bo_stock-1"))
	db.Create(&bl)
	db.Close()
	return dto.NewResult(200, "借书成功", nil)
}
func ReturnBookService(id1, id2 int) dto.Result {
	var book entity.Book
	var bl entity.BookLent
	book.ID = uint(id2)
	bl.StuId = id1
	bl.BookId = id2
	db := utils.GetConn()
	db.Find(&book)
	db.Model(&book).Update("bo_stock", gorm.Expr("bo_stock+1"))
	db.Where("stu_id = ? and book_id = ?", bl.StuId, bl.BookId).Delete(&entity.BookLent{})
	db.Close()
	return dto.NewResult(200, "还书成功", nil)
}

func ShowBooksService() dto.Result {
	db := utils.GetConn()
	var books []entity.Book
	db.Find(&books)
	db.Close()
	return dto.NewResult(200, "请求成功", books)
}

func ShowBookLentService() dto.Result {
	db := utils.GetConn()
	var bl []entity.BookLent
	db.Find(&bl)
	//fmt.Println(bl)
	type Record struct {
		Id        uint    `json:"id"`
		StuName   string  `json:"stu_name,omitempty"`
		StuClass  string  `json:"stu_class,omitempty"`
		BookName  string  `json:"book_name,omitempty"`
		BookPrice float64 `json:"book_price,omitempty"`
	}

	//records := make(map[entity.Student]entity.Book, 10)
	var records []Record
	for _, r := range bl {
		var record Record
		db.Raw("select students.id ,students.stu_name,students.stu_class,books.bo_name as book_name,books.bo_price as book_price from students join books on students.id = ? and books.id = ?", r.StuId, r.BookId).Scan(&record)
		records = append(records, record)

	}
	return dto.NewResult(200, "请求成功", records)
}
