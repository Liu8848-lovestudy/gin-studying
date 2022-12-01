package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func BookManageController(server *gin.Engine) {
	//图书管理
	bookGroup := server.Group("/books")
	//添加图书
	bookGroup.POST("/book", api.AddBook())
	//删除图书
	bookGroup.DELETE("/:id", api.DeleteBook())

	//修改图书
	bookGroup.PUT("/new", api.ModifyBook())

}
