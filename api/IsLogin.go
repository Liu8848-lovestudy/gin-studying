package api

import (
	"gin-studying/dto"
	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var res dto.Result
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			res = dto.NewResult(400, "未登录，请先登录", nil)
			context.JSON(res.Status, res)
			context.Abort()
			return
		}
		res = dto.NewResult(200, "已登录", nil)
		context.JSON(res.Status, res)
	}
}
