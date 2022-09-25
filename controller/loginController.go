package controller

import (
	"github.com/gin-gonic/gin"
	"kuikui/middleware"
	"net/http"
)

type Login struct {
}

func (l Login) LoginController(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if !(username == "admin" && password == "123456") {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	//生成token
	jwt := middleware.NewJWT()
	userInfo := middleware.BaseClaims{}
	userInfo.Username = username
	userInfo.Password = password

	claims := jwt.CreateClaims(userInfo)

	token, err := jwt.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})

}
