package router

import (
	"micro_product/modules/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, ctrlProduct controllers.CtrlProduct) {
	r.POST("/product", ctrlProduct.Create)
	r.GET("/product", ctrlProduct.Show)
	r.GET("/product/:id", ctrlProduct.FindById)
	r.PUT("/products/:id", ctrlProduct.UpdateStock)
	r.PUT("/product/:id", ctrlProduct.Update)
	r.DELETE("/product/:id", ctrlProduct.Delete)

}
