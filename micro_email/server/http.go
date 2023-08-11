package server

import (
	"micro_email/modules/router"
	"micro_email/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuthZ())
	ctrl := NewSetup()
	router.NewRouterEmail(s, ctrl.CtrlEmail)
	s.Run(":6060")
}
