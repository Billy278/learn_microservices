package controllers

import "github.com/gin-gonic/gin"

type CtrlBalance interface {
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	FindByid(ctx *gin.Context)
	FindByidUser(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateByServer(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
