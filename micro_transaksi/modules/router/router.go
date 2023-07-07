package router

import (
	"micro_transaksi/modules/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, ctrlTransaksi controllers.CtrlTransaksi) {
	r.POST("/transaksi", ctrlTransaksi.Transaksi)
	r.GET("/transaksi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "Server berjalan",
		})
	})
}
