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
	db := db.NewDBPostges()
	validate := validator.New()
	repoTransaksi := repository.NewRepoTransaksiImpl(db)
	srvTransaksi := services.NewServTransaksiImpl(repoTransaksi)
	ctrlTransaksi := controllers.NewCtrlTransaksiImpl(srvTransaksi, validate)
	return Controllers{
		CtrlTransaksi: ctrlTransaksi,
	}
}
