package server

import (
	"micro_product/db"
	"micro_product/modules/controllers"
	"micro_product/modules/handler"
	"micro_product/modules/repository"
	"micro_product/modules/services"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	CtrlProduct controllers.CtrlProduct
}

func NewSetup() Controllers {
	db := db.NewDBPostges()
	validate := validator.New()
	repoProduct := repository.NewRepoProductImpl(db)
	srvProduct := services.NewSrvProductImpl(repoProduct)
	ctrlProduct := controllers.NewCtrlProductImpl(srvProduct, validate)
	return Controllers{
		CtrlProduct: ctrlProduct,
	}
}

func GRPCSetup() *handler.HandlerProductImpl {
	db := db.NewDBPostges()
	productRepo := repository.NewRepoProductImpl(db)
	productServ := services.NewSrvProductImpl(productRepo)
	productCtrl := handler.NewHandlerProductImpl(productServ)

	return productCtrl

}
