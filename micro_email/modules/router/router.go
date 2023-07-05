package router

import (
	"micro_email/modules/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouterEmail(r *gin.Engine, ctrlEmail controllers.CtrlEmail) {
	r.POST("/email", ctrlEmail.SendEmail)

}
