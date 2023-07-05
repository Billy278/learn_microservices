package services

import (
	"context"
	"fmt"
	"log"
	"micro_transaksi/modules/model"
	"micro_transaksi/modules/repository"
	"time"
)

type ServTransaksiImpl struct {
	RepoTransaksi repository.RepoTransaksi
}

func NewServTransaksiImpl(repotransaksi repository.RepoTransaksi) ServTransaksi {
	return &ServTransaksiImpl{
		RepoTransaksi: repotransaksi,
	}
}
func (serv *ServTransaksiImpl) SrvCreate(ctx context.Context, transaksiIn model.Transaksi) (err error) {
	logCtx := fmt.Sprintf("%T, SrvCreate", serv)
	log.Println(logCtx)
	dt := time.Now()
	transaksiIn.Tgl_Transaksi = &dt
	err = serv.RepoTransaksi.RepoCreate(ctx, transaksiIn)
	if err != nil {
		return
	}
	return

}
