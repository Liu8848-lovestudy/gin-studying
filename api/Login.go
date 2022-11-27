package api

import (
	"fmt"
	"gin-studying/entity"
	"gin-studying/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var manager entity.Manager
		if err := context.ShouldBind(&manager); err != nil {
			fmt.Println("获取管理员对象失败")
			context.JSON(http.StatusOK, "登录失败")
			log.Fatal(err.Error())
			return
		}
		result := service.LoginService(manager)
		context.JSON(result.Status, result)
	}
}
