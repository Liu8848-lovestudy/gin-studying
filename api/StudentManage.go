package api

import (
	"fmt"
	"gin-studying/dto"
	"gin-studying/entity"
	"gin-studying/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func AddStu() gin.HandlerFunc {
	return func(context *gin.Context) {
		var stu entity.Student
		if err := context.ShouldBind(&stu); err != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "添加失败", nil))
			fmt.Println("添加学生失败")
			log.Fatal(err.Error())
			return
		}
		result := service.AddStudentService(stu)
		context.JSON(result.Status, result)
	}
}

func DeleteStu() gin.HandlerFunc {
	return func(context *gin.Context) {

		id := context.Param("id")
		id1, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "删除学生失败", nil))
			fmt.Println("输入的学生ID有误")
		}
		result := service.DeleteStudentService(id1)
		context.JSON(result.Status, result)
	}
}

func ModifyStudentInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		var stu entity.Student
		if err := context.ShouldBind(&stu); err != nil {
			context.JSON(http.StatusBadRequest, dto.NewResult(400, "修改失败", nil))
			fmt.Println("修改学生信息失败")
			log.Fatal(err.Error())
			return
		}
		result := service.ModifyStudentInfoService(stu)
		context.JSON(result.Status, result)
	}
}
