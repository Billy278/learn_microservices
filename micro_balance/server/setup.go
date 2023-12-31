package server

import (
	"micro_balance/db"
	"micro_balance/modules/controllers"
	"micro_balance/modules/handler"
	"micro_balance/modules/repository"
	"micro_balance/modules/services"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	CtrlBalance controllers.CtrlBalance
}

func NewSetup() Controllers {
	db := db.NewDBPostges()
	validate := validator.New()
	repoBalance := repository.NewRepoBalanceImpl(db)
	servBalance := services.NewServBalanceImpl(repoBalance)
	ctrlBalance := controllers.NewCtrlBalanceImpl(servBalance, validate)
	return Controllers{
		CtrlBalance: ctrlBalance,
	}
}

func GRPCSetup() *handler.GRPCBalanceImpl {
	db := db.NewDBPostges()
	balanceRepo := repository.NewRepoBalanceImpl(db)
	balanceServ := services.NewServBalanceImpl(balanceRepo)
	balanceCtrl := handler.NewGRPCBalanceImpl(balanceServ)

	return balanceCtrl

}
