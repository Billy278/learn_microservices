package server

import (
	"micro_transaksi/db"
	"micro_transaksi/modules/controllers"
	"micro_transaksi/modules/repository"
	"micro_transaksi/modules/services"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	CtrlTransaksi controllers.CtrlTransaksi
}

func NewSetup() Controllers {

	dbStore := db.NewDBPostges()
	ch := db.NewRabbitMQ()
	validate := validator.New()
	repoTransaksi := repository.NewRepoTransaksiImpl(dbStore)
	srvTransaksi := services.NewServTransaksiImpl(repoTransaksi)
	ctrlTransaksi := controllers.NewCtrlTransaksiImpl(srvTransaksi, validate, ch)
	return Controllers{
		CtrlTransaksi: ctrlTransaksi,
	}
}
