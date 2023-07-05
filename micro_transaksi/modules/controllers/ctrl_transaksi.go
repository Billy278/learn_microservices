package controllers

import "github.com/gin-gonic/gin"

type CtrlTransaksi interface {
	Transaksi(ctx *gin.Context)
}
