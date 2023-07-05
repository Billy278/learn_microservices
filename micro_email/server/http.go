package server

import (
	"micro_email/modules/router"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger())
	ctrl := NewSetup()
	router.NewRouterEmail(s, ctrl.CtrlEmail)
	s.Run(":6060")
}
