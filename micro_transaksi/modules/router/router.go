package router

import (
	"micro_transaksi/modules/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, ctrlTransaksi controllers.CtrlTransaksi) {
	r.POST("/transaksi", ctrlTransaksi.Transaksi)
}
