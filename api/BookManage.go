package api

import (
	"fmt"
	"gin-studying/dto"
	"gin-studying/entity"
	"gin-studying/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func AddBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		var book entity.Book
		if err := context.ShouldBind(&book); err != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "图书添加失败", nil))
			fmt.Println("添加图书失败")
			log.Fatal(err.Error())
			return
		}
		result := service.AddBookService(book)
		context.JSON(result.Status, result)
	}
}

func DeleteBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		id1, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "图书删除失败", nil))
			fmt.Println("输入的书籍ID有误")
		}
		result := service.DeleteBookService(id1)
		context.JSON(result.Status, result)
	}
}
func ModifyBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		var book entity.Book
		if err := context.ShouldBind(&book); err != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "图书信息修改失败", nil))
			fmt.Println("修改图书信息失败")
			log.Fatal(err.Error())
			return
		}
		result := service.ModifyBookService(book)
		context.JSON(result.Status, result)
	}
}
