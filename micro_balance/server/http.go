package server

import (
	"micro_balance/modules/router"
	"micro_balance/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuth())
	ctrl := NewSetup()
	router.NewRouterBlnc(s, ctrl.CtrlBalance)
	s.Run(":8090")
}
