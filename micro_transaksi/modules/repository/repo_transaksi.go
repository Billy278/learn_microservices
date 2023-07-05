package repository

import (
	"context"
	"micro_transaksi/modules/model"
)

type RepoTransaksi interface {
	RepoCreate(ctx context.Context, transaksiIn model.Transaksi) (err error)
}
