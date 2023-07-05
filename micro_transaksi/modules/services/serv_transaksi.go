package services

import (
	"context"
	"micro_transaksi/modules/model"
)

type ServTransaksi interface {
	SrvCreate(ctx context.Context, transaksiIn model.Transaksi) (err error)
}
