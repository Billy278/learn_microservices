package repository

import (
	"context"
	"micro_balance/modules/model"
)

type RepoBalance interface {
	RepoShow(ctx context.Context) (balanceRes []model.Balance, err error)
	RepoCreate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error)
	RepoFinById(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error)
	RepoFinByIdUser(ctx context.Context, userId uint64) (balanceRes model.Balance, err error)
	RepoUpdateByUser(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error)
	RepoUpdate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error)
	RepoDelete(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error)
}
