package routes

import (
	"github.com/gin-gonic/gin"
	"kuikui/controller"
)

func LoginRouteInit(r *gin.Engine) {
	r.POST("/login", controller.Login{}.LoginController)
}
