package controllers

import (
	"errors"
	"fmt"
	"micro_product/modules/model"
	"micro_product/modules/services"
	"micro_product/pkg/crypto"
	"micro_product/pkg/middleware"
	"micro_product/pkg/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlProductImpl struct {
	ServiceProduct services.SrvProduct
	Validate       *validator.Validate
}

func NewCtrlProductImpl(serviceproduct services.SrvProduct, validate *validator.Validate) CtrlProduct {
	return &CtrlProductImpl{
		ServiceProduct: serviceproduct,
		Validate:       validate,
	}
}

func (ctrl *CtrlProductImpl) Show(ctx *gin.Context) {
	res, err := ctrl.ServiceProduct.SrvShow(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, responses.FailRes{
			Code:    http.StatusBadGateway,
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
func (ctrl *CtrlProductImpl) Create(ctx *gin.Context) {
	reqProduct := model.Product{}
	if err := ctx.ShouldBindJSON(&reqProduct); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	//validasi request
	err := ctrl.Validate.Struct(reqProduct)
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
	id, _ := strconv.Atoi(accessClaim.UserId)
	reqProduct.UserId = uint64(id)
	res, err := ctrl.ServiceProduct.SrvCreate(ctx, reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.InternalServer,
			Error:   err.Error(),
		})
		return
	}
	message := fmt.Sprintf("product %v success create", res.Name)
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: message,
	})
}
func (ctrl *CtrlProductImpl) FindById(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}

	res, err := ctrl.ServiceProduct.SrvFindBy(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.FailRes{
			Code:    http.StatusNotFound,
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
func (ctrl *CtrlProductImpl) Update(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}

	reqproduct := model.Product{}
	err = ctx.ShouldBindJSON(&reqproduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//validasi request
	err = ctrl.Validate.Struct(reqproduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}

	reqproduct.Id = id
	res, err := ctrl.ServiceProduct.SrvUpdate(ctx, reqproduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	message := fmt.Sprintf("product id=%v success update ", res.Id)
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: message,
	})
}
func (ctrl *CtrlProductImpl) Delete(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}
	res, err := ctrl.ServiceProduct.SrvDelete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	message := fmt.Sprintf("product id=%v success Delete ", res.Id)
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: message,
	})
}

func (ctrl *CtrlProductImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
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

func (ctrl *CtrlProductImpl) UpdateStock(ctx *gin.Context) {
	id, err := ctrl.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidParam,
			Error:   err.Error(),
		})
		return
	}

	reqproduct := model.Product{}
	err = ctx.ShouldBindJSON(&reqproduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	reqproduct.Id = id
	res, err := ctrl.ServiceProduct.SrvUpdateStock(ctx, reqproduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	message := fmt.Sprintf("product id=%v success update ", res.Id)
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: message,
	})
}
