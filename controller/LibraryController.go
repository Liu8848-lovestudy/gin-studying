package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func LibraryCon(server *gin.Engine) {
	//图书借阅
	server.POST("/addBookLent/:stuId/:bookId", api.LendBook())
	//图书归还
	server.DELETE("/deleteBookLent/:stuId/:bookId", api.ReturnBook())

	//图书信息展示
	server.GET("/showBooks", api.ShowBooks())

	//图书借出列表展示
	server.GET("/showBookLent", api.ShowBookLent())
}
