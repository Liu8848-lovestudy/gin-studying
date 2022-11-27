package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func BookManageController(server *gin.Engine) {
	//图书管理
	bookGroup := server.Group("/book")
	//添加图书
	bookGroup.POST("/add", api.AddBook())
	//删除图书
	bookGroup.DELETE("/delete/:id", api.DeleteBook())

	//修改图书
	bookGroup.POST("/modify", api.ModifyBook())

}
