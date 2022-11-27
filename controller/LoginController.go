package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func LoginCon(server *gin.Engine) {
	server.POST("/login", api.Login())
}
