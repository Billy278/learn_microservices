package router

import (
	"micro_balance/modules/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouterBlnc(r *gin.Engine, ctrlBalance controllers.CtrlBalance) {
	r.GET("/balance", ctrlBalance.Show)
	r.POST("/balance", ctrlBalance.Create)
	r.GET("/balance/:id", ctrlBalance.FindByidUser)
	r.GET("/balances/:id", ctrlBalance.FindByid)
	r.PUT("/balances/:id", ctrlBalance.Update)
	r.PUT("/balance/:id", ctrlBalance.UpdateByServer)
	r.DELETE("/balance/:id", ctrlBalance.Delete)

}
