package api

import (
	"gin-studying/dto"
	"gin-studying/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LendBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		stuId := context.Param("stuId")
		bookId := context.Param("bookId")
		id1, err1 := strconv.Atoi(stuId)
		id2, err2 := strconv.Atoi(bookId)
		if err1 != nil && err2 != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "借书失败", nil))
			panic("输入参数有误，借书失败")
		}
		result := service.LendBookService(id1, id2)
		context.JSON(result.Status, result)
	}
}
func ReturnBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		stuId := context.Param("stuId")
		bookId := context.Param("bookId")
		id1, err1 := strconv.Atoi(stuId)
		id2, err2 := strconv.Atoi(bookId)
		if err1 != nil && err2 != nil {
			context.JSON(http.StatusOK, dto.NewResult(400, "还书失败", nil))
			panic("输入参数有误，还书失败")
		}
		result := service.ReturnBookService(id1, id2)
		context.JSON(result.Status, result)
	}
}
func ShowBooks() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := service.ShowBooksService()
		context.JSON(result.Status, result)
	}
}
func ShowBookLent() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := service.ShowBookLentService()
		context.JSON(result.Status, result)
	}
}
