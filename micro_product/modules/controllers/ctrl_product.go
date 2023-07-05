package controllers

import "github.com/gin-gonic/gin"

type CtrlProduct interface {
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UpdateStock(ctx *gin.Context)
}
