package service

import (
	"gin-studying/dto"
	"gin-studying/entity"
	"gin-studying/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func LoginService(manager entity.Manager) dto.Result {
	db := utils.GetConn()
	var m entity.Manager
	db.Where("user_name=?", manager.UserName).Find(&m)
	if m.PassWord == manager.PassWord {
		expireTime := time.Now().Add(30 * time.Minute) //设置过期时长
		manager.StandardClaims = jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //开始生效时间
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //颁发时间
			Issuer:    "localhost",       //颁发者
			Subject:   "manager token",   //签名主题
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, manager)
		//fmt.Println(token)
		tokenString, err := token.SignedString([]byte("login"))
		if err != nil {
			panic(err)
		}
		//fmt.Println(tokenString)
		return dto.NewResult(200, "登录成功", gin.H{"token": tokenString})

	} else {
		return dto.NewResult(400, "账号或密码错误", nil)
	}
}
