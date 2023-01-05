package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func GetConn() *gorm.DB {
	//str := "root:admin123@tcp(127.0.0.1:3306)/library?charset=utf8&parseTime=true&loc=Local"
	str := "root:admin123@tcp(mysql01:3306)/ginLibrary?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", str)
	if err != nil {
		fmt.Println("数据库连接出错")
		log.Fatal(err.Error())
	}
	return db
}
