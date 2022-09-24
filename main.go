package main

import (
	"github.com/gin-gonic/gin"
	"kuikui/middleware"
	"kuikui/routes"
)

func main() {
	//middleware.CasbinTest()

	r := gin.Default()
	r.Use(middleware.CasbinHandler())

	routes.SysRoutesInit(r)

	err := r.Run(":9000")

	if err != nil {
		//fmt.Print(err.Error)
		panic("系统起动失败:" + err.Error())
	}
}
