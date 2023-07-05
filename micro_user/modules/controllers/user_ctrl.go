package controllers

import "github.com/gin-gonic/gin"

type UserCtrl interface {
	LoginUser(ctx *gin.Context)
	Register(ctx *gin.Context)
	Home(ctx *gin.Context)
}
