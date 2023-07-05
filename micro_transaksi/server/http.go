package server

import (
	"micro_transaksi/modules/router"
	"micro_transaksi/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuth())

	ctrl := NewSetup()
	router.NewRouter(s, ctrl.CtrlTransaksi)
	s.Run(":7070")

}
