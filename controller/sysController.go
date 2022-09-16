package controller

import (
	"github.com/gin-gonic/gin"
	"kuikui/models"
	"net/http"
)

type SysController struct {
}

func (sys SysController) User(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "很多很多用户",
	})
}

func (sys SysController) CreateTable(ctx *gin.Context) {
	err1 := models.DB.AutoMigrate(&models.User{})
	if err1 != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "很多很多用户",
	})
}
