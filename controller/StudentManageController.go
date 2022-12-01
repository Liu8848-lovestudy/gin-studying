package controller

import (
	"gin-studying/api"
	"github.com/gin-gonic/gin"
)

func StudentManageController(server *gin.Engine) {
	stuGroup := server.Group("/students")

	stuGroup.Use(api.IsLogin())
	//添加学生
	stuGroup.POST("/student", api.AddStu())
	//删除学生
	stuGroup.DELETE("/:id")
	//修改学生信息
	stuGroup.PUT("/new", api.ModifyStudentInfo())

}
