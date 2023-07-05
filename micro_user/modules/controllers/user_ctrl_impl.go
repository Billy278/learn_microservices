package controllers

import (
	"fmt"
	"micro_user/modules/model"
	"micro_user/modules/services"
	"micro_user/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlImpl struct {
	UserServ services.UserServ
	Validate validator.Validate
}

func NewUserCtrlImpl(userserv services.UserServ, validate validator.Validate) UserCtrl {
	return &UserCtrlImpl{
		UserServ: userserv,
		Validate: validate,
	}
}

func (user_ctrl *UserCtrlImpl) LoginUser(ctx *gin.Context) {
	userReq := model.Login{}
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	//validasi request

	err := user_ctrl.Validate.Struct(userReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	tokens, err := user_ctrl.UserServ.ServLoginUser(ctx, userReq.Username, userReq.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: tokens,
	})
}
func (user_ctrl *UserCtrlImpl) Register(ctx *gin.Context) {
	userReq := model.User{}
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := user_ctrl.Validate.Struct(userReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	userRes, err := user_ctrl.UserServ.ServCreateUser(ctx, userReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, responses.SuccessRes{
		Code: http.StatusAccepted,
		Data: userRes,
	})
}
func (user_ctrl *UserCtrlImpl) Home(ctx *gin.Context) {
	fmt.Fprint(ctx.Writer, "<h1>halaman home</h1>")
	ctx.Writer.Header().Set("Content-type", "text/html")
	ctx.JSON(http.StatusUnauthorized, responses.SuccessRes{
		Code: http.StatusOK,
	})

}
