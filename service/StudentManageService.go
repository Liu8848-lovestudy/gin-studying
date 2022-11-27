package service

import (
	"gin-studying/dto"
	"gin-studying/entity"
	"gin-studying/utils"
)

func AddStudentService(stu entity.Student) dto.Result {
	db := utils.GetConn()
	db.Create(&stu)
	db.Close()
	return dto.NewResult(200, "添加成功", stu)
}

func DeleteStudentService(id1 int) dto.Result {
	var stu entity.Student
	stu.ID = uint(id1)
	db := utils.GetConn()
	db.Debug().Delete(&stu)
	db.Close()
	return dto.NewResult(200, "删除成功", nil)
}
func ModifyStudentInfoService(stu entity.Student) dto.Result {
	db := utils.GetConn()
	db.Model(&entity.Student{}).Update(&stu)
	db.Close()
	return dto.NewResult(200, "学生信息修改成功", nil)
}
