package routes

import (
	"github.com/gin-gonic/gin"
	"kuikui/controller"
)

func SysRoutesInit(router *gin.Engine) {
	r := router.Group("/sys")
	{
		r.GET("/user", controller.SysController{}.User)
		r.GET("/createtable", controller.SysController{}.CreateTable)
	}
}
