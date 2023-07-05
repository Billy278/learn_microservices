package services

import (
	"context"
	"micro_balance/modules/model"
)

type ServBalance interface {
	ServShow(ctx context.Context) (balanceRes []model.Balance, err error)
	ServCreate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error)
	ServFinById(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error)
	ServFinByIdUser(ctx context.Context, userId uint64) (balanceRes model.Balance, err error)
	ServUpdate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error)
	ServUpdateByIdUser(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error)
	ServDelete(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error)
}
