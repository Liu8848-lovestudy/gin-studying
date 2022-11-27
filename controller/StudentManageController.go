package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func StudentManageController(server *gin.Engine) {
	stuGroup := server.Group("/student")

	stuGroup.Use(api.IsLogin())
	//添加学生
	stuGroup.POST("/add", api.AddStu())
	//删除学生
	stuGroup.DELETE("/delete/:id")
	//修改学生信息
	stuGroup.POST("/modify", api.ModifyStudentInfo())

}
