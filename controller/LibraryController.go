package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func LibraryCon(server *gin.Engine) {
	libraryGroup := server.Group("/library")
	//图书借阅
	libraryGroup.POST("/book/:stuId/:bookId", api.LendBook())
	//图书归还
	libraryGroup.DELETE("/:stuId/:bookId", api.ReturnBook())
	//图书信息展示
	libraryGroup.GET("/books", api.ShowBooks())
	//图书借出列表展示
	libraryGroup.GET("/bookLents", api.ShowBookLent())
}
