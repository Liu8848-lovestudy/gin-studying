package controller

import "github.com/gin-gonic/gin"

func StartServer() {
	server := gin.Default()
	LoginCon(server)
	LibraryCon(server)
	BookManageController(server)
	StudentManageController(server)
	server.Run(":8080")
}
