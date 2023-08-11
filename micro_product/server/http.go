package server

import (
	"micro_product/modules/router"
	"micro_product/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuthZ())

	ctrl := NewSetup()
	router.NewRouter(s, ctrl.CtrlProduct)
	s.Run(":9090")

}
