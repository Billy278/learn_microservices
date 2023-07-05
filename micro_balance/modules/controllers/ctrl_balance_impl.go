package controllers

import (
	"errors"
	"fmt"
	"micro_balance/modules/model"
	"micro_balance/modules/services"
	"micro_balance/pkg/crypto"
	"micro_balance/pkg/middleware"
	"micro_balance/pkg/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlBalanceImpl struct {
	BalanceServ services.ServBalance
	Validate    *validator.Validate
}

func NewCtrlBalanceImpl(balanceserv services.ServBalance, validate *validator.Validate) CtrlBalance {
	return &CtrlBalanceImpl{
		BalanceServ: balanceserv,
		Validate:    validate,
	}
}

func (ctrl *CtrlBalanceImpl) Show(ctx *gin.Context) {
	res, err := ctrl.BalanceServ.ServShow(ctx)
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
		Data: res,
	})
}
func (ctrl *CtrlBalanceImpl) Create(ctx *gin.Context) {
	balanceReq := model.Balance{}
	err := ctx.ShouldBindJSON(&balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request

	err = ctrl.Validate.Struct(balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//get id user from token

	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		fmt.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	var accessClaim model.AccessClaim

	if err := crypto.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		fmt.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidPayload,
			Error:   err.Error(),
		})
		return
	}
	usrId, err := strconv.Atoi(accessClaim.UserId)
	if err != nil {
		fmt.Printf("[ERROR] Get token usr id:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	balanceReq.UserId = uint64(usrId)
	res, err := ctrl.BalanceServ.ServCreate(ctx, balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	data := fmt.Sprintf("Sucses create akun balance with user id =%v", res.UserId)
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: data,
	})

}
func (ctrl *CtrlBalanceImpl) FindByid(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}
	res, err := ctrl.BalanceServ.ServFinById(ctx, id)
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
		Data: res,
	})
}
func (ctrl *CtrlBalanceImpl) UpdateByServer(ctx *gin.Context) {
	balanceReq := model.Balance{}
	err := ctx.ShouldBindJSON(&balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	// validasi request
	err = ctrl.Validate.Struct(balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}

	//get iduser from token

	balanceReq.UserId = id
	res, err := ctrl.BalanceServ.ServUpdateByIdUser(ctx, balanceReq)
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
		Data: res,
	})
}
func (ctrl *CtrlBalanceImpl) Update(ctx *gin.Context) {
	balanceReq := model.Balance{}
	err := ctx.ShouldBindJSON(&balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	// validasi request
	err = ctrl.Validate.Struct(balanceReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}

	//get iduser from token

	accessClaimIn, ok := ctx.Get(string(middleware.AccessClaim))
	if !ok {
		err := errors.New("error get claim from context")
		fmt.Printf("[ERROR] Get Payload:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	var accessClaim model.AccessClaim

	if err := crypto.ObjectMapper(accessClaimIn, &accessClaim); err != nil {
		fmt.Printf("[ERROR] Get claim from context:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidPayload,
			Error:   err.Error(),
		})
		return
	}
	usrId, err := strconv.Atoi(accessClaim.UserId)
	if err != nil {
		fmt.Printf("[ERROR] Get token usr id:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	balanceReq.Id = id
	balanceReq.UserId = uint64(usrId)
	res, err := ctrl.BalanceServ.ServUpdate(ctx, balanceReq)
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
		Data: res,
	})
}
func (ctrl *CtrlBalanceImpl) Delete(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}
	res, err := ctrl.BalanceServ.ServDelete(ctx, id)
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
		Data: res,
	})
}

func (ctrl *CtrlBalanceImpl) FindByidUser(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}
	res, err := ctrl.BalanceServ.ServFinByIdUser(ctx, id)
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
		Data: res,
	})
}
func (ctrl *CtrlBalanceImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		return
	}

	return idUint, err

}
