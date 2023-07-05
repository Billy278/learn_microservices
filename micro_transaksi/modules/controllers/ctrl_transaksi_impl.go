package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"micro_transaksi/modules/model"
	"micro_transaksi/modules/services"
	"micro_transaksi/pkg/crypto"
	"micro_transaksi/pkg/middleware"
	"micro_transaksi/pkg/responses"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rabbitmq/amqp091-go"
)

type CtrlTransaksiImpl struct {
	ServTransaksi services.ServTransaksi
	Validate      *validator.Validate
	Ch            *amqp091.Channel
}

func NewCtrlTransaksiImpl(servtransaksi services.ServTransaksi, validate *validator.Validate, ch *amqp091.Channel) CtrlTransaksi {
	return &CtrlTransaksiImpl{
		ServTransaksi: servtransaksi,
		Validate:      validate,
		Ch:            ch,
	}
}

func (ctrl *CtrlTransaksiImpl) Transaksi(ctx *gin.Context) {
	transaksiReq := model.Transaksi{}
	err := ctx.ShouldBindJSON(&transaksiReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//validasi request
	err = ctrl.Validate.Struct(transaksiReq)
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

	iduser, err := strconv.Atoi(accessClaim.UserId)
	if err != nil {
		fmt.Printf("[ERROR] Get userid:%v\n", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	transaksiReq.Id_User = uint64(iduser)

	//get data product by id

	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: http.DefaultTransport,
	}
	//===================================================================================
	//get data user
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/product/"+strconv.Itoa(int(transaksiReq.Id_Product)), nil)
	//req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/product/10", nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}

	req.Header.Add("KEY", crypto.SharedKey)
	resGetProduct, err := client.Do(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}

	data, err := ioutil.ReadAll(resGetProduct.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}

	resData := responses.SuccessRes{}
	err = json.Unmarshal(data, &resData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}

	jsonData, _ := json.Marshal(resData.Data)
	product := model.Product{}
	json.Unmarshal(jsonData, &product)
	if product.Stock < transaksiReq.Quantity {

		err = errors.New("STOCK tidak cukup")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return

	} else {
		product.Stock = product.Stock - transaksiReq.Quantity
	}

	//===================================================================================
	//get data balance
	req, err = http.NewRequest(http.MethodGet, "http://localhost:8090/balance/"+strconv.Itoa(int(transaksiReq.Id_User)), nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	req.Header.Add("KEY", crypto.SharedKey)
	resGetBalance, err := client.Do(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	dataBalance, err := ioutil.ReadAll(resGetBalance.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return

	}

	err = json.Unmarshal(dataBalance, &resData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}

	resBalance, _ := json.Marshal(resData.Data)
	balance := model.Balance{}
	json.Unmarshal(resBalance, &balance)
	//	cek total biasa dan saldo user
	total := float64(transaksiReq.Quantity) * product.Harga
	fmt.Println(total)
	if total > balance.Saldo {
		err = errors.New("SALDO TIDAK CUKUP")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.FailRes{
			Code:    http.StatusBadRequest,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	} else {
		balance.Saldo = balance.Saldo - total
	}
	transaksiReq.Total = total
	//===================================================================================
	//send data ke gmail services
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		bodyEm := fmt.Sprintf(`{
			"name_receiver":"%v",
			"name_product" :"%v",
			"harga":%v,
			"qty":%v,
			"total":%v,
			"email_receiver":"%v"}`, accessClaim.Username, product.Name, product.Harga, transaksiReq.Quantity, total, accessClaim.Email)
		err := ctrl.Ch.PublishWithContext(
			ctx,
			"ex_learn_micro",
			"PWSD",
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/json",
				Body:        []byte(bodyEm),
			})
		if err != nil {
			panic(err)
		}
		//without message broker
		// bodyEm := fmt.Sprintf(`{
		// 	"name_receiver":"%v",
		// 	"name_product" :"%v",
		// 	"harga":%v,
		// 	"qty":%v,
		// 	"total":%v,
		// 	"email_receiver":"%v"}`, accessClaim.Username, product.Name, product.Harga, transaksiReq.Quantity, total, accessClaim.Email)
		// BodyEmail := strings.NewReader(bodyEm)
		// reqEmail, err := http.NewRequest(http.MethodPost, "http://localhost:6060/email", BodyEmail)
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
		// 		Code:    http.StatusInternalServerError,
		// 		Message: responses.SomethingWentWrong,
		// 		Error:   err.Error(),
		// 	})
		// 	return
		// }
		// client.Do(reqEmail)
	}()

	//jika semua kriteria terpenuhi maka lakukan update data pada table product dan balance
	//=====================================================================================
	//===================================================================================
	//update data product
	bodypro := fmt.Sprintf(`{"stock":%v}`, product.Stock)
	bodyProduct := strings.NewReader(bodypro)
	requpdateProduct, err := http.NewRequest(http.MethodPut, "http://localhost:9090/products/"+strconv.Itoa(int(transaksiReq.Id_Product)), bodyProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	requpdateProduct.Header.Add("KEY", crypto.SharedKey)
	_, err = client.Do(requpdateProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	//===================================================================================
	//udpate data balance
	bodyblc := fmt.Sprintf(`{"saldo":%v}`, balance.Saldo)
	fmt.Println("aaaaaaa")
	fmt.Println(balance.Saldo)
	bodyBalance := strings.NewReader(bodyblc)
	reqBalanceUpdate, err := http.NewRequest(http.MethodPut, "http://localhost:8090/balance/"+strconv.Itoa(int(transaksiReq.Id_User)), bodyBalance)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	reqBalanceUpdate.Header.Add("KEY", crypto.SharedKey)
	_, err = client.Do(reqBalanceUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	err = ctrl.ServTransaksi.SrvCreate(ctx, transaksiReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.FailRes{
			Code:    http.StatusInternalServerError,
			Message: responses.SomethingWentWrong,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, responses.SuccessRes{
		Code: http.StatusOK,
		Data: "Transaksi Sussces",
	})
	fmt.Println(accessClaim.Email)
	wg.Wait()
}

// func (ctrl *CtrlTransaksiImpl) convertStringToUint64(id_product, qty string) (idproductRes, qtyRes uint64, err error) {

// 	// transform id string to uint64
// 	idproductRes, err = strconv.ParseUint(id_product, 10, 64)
// 	if err != nil {
// 		err = errors.New("failed parse id")
// 		return
// 	}

// 	qtyRes, err = strconv.ParseUint(qty, 10, 64)
// 	if err != nil {
// 		err = errors.New("failed parse id")
// 		return
// 	}
// 	return
// }
