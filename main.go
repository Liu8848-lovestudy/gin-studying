package main

import (
	"fmt"
	"gin-studying/entity"
	"gin-studying/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

var flag bool

func isLogin() gin.HandlerFunc {
	//context.Set("status","已登录")
	return func(context *gin.Context) {
		//value, exists := context.Get("status")
		//fmt.Println(value, exists)
		if flag {
			context.JSON(200, "已登录")
		} else {
			context.JSON(200, "未登录，请先登录")
			context.Abort()
		}
	}
}

func main() {
	//在数据库中创建表
	//utils.CreateTables()
	server := gin.Default()
	//登录
	server.POST("/login", func(context *gin.Context) {
		var manager entity.Manager
		if err := context.ShouldBind(&manager); err != nil {
			fmt.Println("获得图书馆对象失败")
			log.Fatal(err.Error())
			return
		}
		db := utils.GetConn()
		var m entity.Manager
		db.Where("user_name=?", manager.UserName).Find(&m)
		if m.PassWord == manager.PassWord {
			//context.Set("status", "已登录")
			flag = true
			context.JSON(200, "登录成功")
		} else {
			context.JSON(200, "账号或密码错误")
		}
	})

	//图书借阅
	server.GET("/borrow/:stuId/:bookId", func(context *gin.Context) {
		//代码
		stuId := context.Param("stuId")
		bookId := context.Param("bookId")
		var bl entity.BookLent
		id1, err1 := strconv.Atoi(stuId)
		id2, err2 := strconv.Atoi(bookId)
		if err1 != nil && err2 != nil {
			panic("输入参数有误，借书失败")
		}
		var book entity.Book
		book.ID = uint(id2)
		bl.StuId = id1
		bl.BookId = id2
		db := utils.GetConn()
		db.Find(&book)
		if book.BoStock <= 0 {
			context.JSON(200, "库存不足")
			return
		}
		db.Model(&book).Update("bo_stock", gorm.Expr("bo_stock-1"))
		db.Create(&bl)
		db.Close()
		context.JSON(http.StatusOK, "借书成功")
	})
	//图书归还
	server.GET("/return/:stuId/:bookId", func(context *gin.Context) {
		//代码
		stuId := context.Param("stuId")
		bookId := context.Param("bookId")
		var bl entity.BookLent
		id1, err1 := strconv.Atoi(stuId)
		id2, err2 := strconv.Atoi(bookId)
		if err1 != nil && err2 != nil {
			panic("输入参数有误，还书失败")
		}
		var book entity.Book
		book.ID = uint(id2)
		bl.StuId = id1
		bl.BookId = id2
		db := utils.GetConn()
		db.Find(&book)
		db.Model(&book).Update("bo_stock", gorm.Expr("bo_stock+1"))
		db.Delete(&bl)
		db.Close()
		context.JSON(http.StatusOK, "还书成功")
	})
	//图书信息展示
	server.GET("/showBooks", func(context *gin.Context) {
		db := utils.GetConn()
		var books []entity.Book
		db.Find(&books)
		context.JSON(http.StatusOK, books)
		db.Close()
	})

	//图书借出列表展示
	server.GET("/showLent", func(context *gin.Context) {
		//代码
		db := utils.GetConn()
		var bl []entity.BookLent
		db.Find(&bl)
		fmt.Println(bl)
		//records := make(map[entity.Student]entity.Book, 10)
		for _, r := range bl {
			var stu entity.Student
			var book entity.Book
			stu.ID = uint(r.StuId)
			book.ID = uint(r.BookId)
			db.Find(&stu)
			db.Find(&book)
			fmt.Println(stu.ID, stu.StuName, book.BoName, book.BoPrice)
			context.JSON(200, stu)
			context.JSON(200, book)
			//records[stu] = book
		}

		//for stu, book := range records {
		//	fmt.Println(stu.ID, stu.StuName, book.BoName, book.BoPrice)
		//	context.JSON(200, stu)
		//	context.JSON(200, book)
		//}
	})

	//学生管理
	stuGroup := server.Group("/student")

	stuGroup.Use(isLogin())
	//添加学生
	stuGroup.POST("/add", func(context *gin.Context) {
		var stu entity.Student
		if err := context.ShouldBind(&stu); err != nil {
			fmt.Println("添加学生失败")
			log.Fatal(err.Error())
			return
		}
		db := utils.GetConn()
		db.Create(&stu)
		context.JSON(http.StatusOK, stu)
		db.Close()
	})
	//删除学生
	stuGroup.DELETE("/delete/:id", func(context *gin.Context) {
		var stu entity.Student
		id := context.Param("id")
		if id == "" {
			fmt.Println("删除学生失败")
			return
		}
		id1, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("输入的学生ID有误")
		}
		stu.ID = uint(id1)
		db := utils.GetConn()
		db.Debug().Delete(&stu)
		context.JSON(http.StatusOK, "删除学生成功")
		db.Close()
	})
	//修改学生信息
	stuGroup.POST("/modify", func(context *gin.Context) {
		var stu entity.Student
		if err := context.ShouldBind(&stu); err != nil {
			fmt.Println("修改学生信息失败")
			log.Fatal(err.Error())
			return
		}
		db := utils.GetConn()
		db.Model(&entity.Student{}).Update(&stu)
		context.JSON(http.StatusOK, "学生信息修改成功")
		db.Close()
	})

	//图书管理
	bookGroup := server.Group("/book")
	//添加图书
	bookGroup.POST("/add", func(context *gin.Context) {
		var book entity.Book
		if err := context.ShouldBind(&book); err != nil {
			fmt.Println("添加图书失败")
			log.Fatal(err.Error())
			return
		}
		db := utils.GetConn()
		db.Create(&book)
		context.JSON(http.StatusOK, book)
		db.Close()
	})
	//删除图书
	bookGroup.DELETE("/delete/:id", func(context *gin.Context) {
		var book entity.Book
		id := context.Param("id")
		if id == "" {
			fmt.Println("删除书籍失败")
			return
		}
		id1, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("输入的书籍ID有误")
		}
		book.ID = uint(id1)
		db := utils.GetConn()
		db.Debug().Delete(&book)
		context.JSON(http.StatusOK, "删除书籍成功")
		db.Close()
	})

	//修改图书
	bookGroup.POST("/modify", func(context *gin.Context) {
		var book entity.Book
		if err := context.ShouldBind(&book); err != nil {
			fmt.Println("修改图书信息失败")
			log.Fatal(err.Error())
			return
		}
		db := utils.GetConn()
		db.Model(&entity.Book{}).Update(&book)
		context.JSON(http.StatusOK, "图书信息修改成功")
		db.Close()
	})
	server.Run(":8080")
}
