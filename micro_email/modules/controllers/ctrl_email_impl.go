package controllers

import (
	"fmt"
	"micro_email/modules/model"
	"micro_email/modules/services"
	"micro_email/pkg/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlEmailImpl struct {
	EmailServ services.ServEmail
	Validate  *validator.Validate
}

func NewCtrlEmailImpl(emailserv services.ServEmail, validate *validator.Validate) CtrlEmail {
	return &CtrlEmailImpl{
		EmailServ: emailserv,
		Validate:  validate,
	}
}
func (ctrl *CtrlEmailImpl) SendEmail(ctx *gin.Context) {
	emailReq := model.Email{}
	err := ctx.ShouldBindJSON(&emailReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//validasi request
	err = ctrl.Validate.Struct(emailReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	res, err := ctrl.EmailServ.ServCreate(ctx, emailReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	rs := fmt.Sprintf("pesan ke %v berhasil di kirim", res.Name_Receiver)
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: rs,
	})
}
