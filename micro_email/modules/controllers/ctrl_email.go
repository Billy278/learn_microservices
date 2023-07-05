package controllers

import "github.com/gin-gonic/gin"

type CtrlEmail interface {
	SendEmail(ctx *gin.Context)
}
