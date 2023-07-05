package server

import (
	routerUser "micro_user/modules/router/v1/user"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())
	ctrl := NewSetup()
	routerUser.NewUserRouter(r, ctrl.UserCtrl)
	r.Run(":8080")
}
