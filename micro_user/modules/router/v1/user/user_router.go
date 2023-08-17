package user

import (
	"micro_user/modules/controllers"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, usrCrl controllers.UserCtrl) {
	r.POST("/register", usrCrl.Register)
	r.POST("/login", usrCrl.LoginUser)

	// r.GET("/userr/home", middleware.TestBearerOAuth(), usrCrl.Home)
	// //grouping
	// user := r.Group("/user", middleware.BearerOAuth())
	// user.GET("/home", usrCrl.Home)

}
