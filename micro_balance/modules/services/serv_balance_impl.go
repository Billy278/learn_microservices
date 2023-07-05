package services

import (
	"context"
	"fmt"
	"micro_balance/modules/model"
	"micro_balance/modules/repository"
)

type ServBalanceImpl struct {
	BalanceRepo repository.RepoBalance
}

func NewServBalanceImpl(balancerepo repository.RepoBalance) ServBalance {
	return &ServBalanceImpl{
		BalanceRepo: balancerepo,
	}
}

func (serv *ServBalanceImpl) ServShow(ctx context.Context) (balanceRes []model.Balance, err error) {
	log := fmt.Sprintf("%T,ServShow", serv)
	fmt.Println(log)
	balanceRes, err = serv.BalanceRepo.RepoShow(ctx)
	if err != nil {
		return
	}
	return

}
func (serv *ServBalanceImpl) ServCreate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,ServCreate", serv)
	fmt.Println(log)
	balanceRes, err = serv.BalanceRepo.RepoCreate(ctx, balanceIn)
	if err != nil {
		return
	}
	return
}
func (serv *ServBalanceImpl) ServFinById(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,ServFinById", serv)
	fmt.Println(log)
	balanceRes, err = serv.BalanceRepo.RepoFinById(ctx, balanceId)
	if err != nil {
		return
	}
	return
}
func (serv *ServBalanceImpl) ServFinByIdUser(ctx context.Context, userId uint64) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,ServFinByIdUser", serv)
	fmt.Println(log)
	balanceRes, err = serv.BalanceRepo.RepoFinByIdUser(ctx, userId)
	if err != nil {
		return
	}
	return
}
func (serv *ServBalanceImpl) ServUpdateByIdUser(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,ervUpdateByIdUser", serv)
	fmt.Println(log)
	_, err = serv.BalanceRepo.RepoFinByIdUser(ctx, balanceIn.UserId)
	if err != nil {
		return
	}
	balanceRes, err = serv.BalanceRepo.RepoUpdateByUser(ctx, balanceIn)
	if err != nil {
		return
	}
	return
}
func (serv *ServBalanceImpl) ServUpdate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,ServUpdate", serv)
	fmt.Println(log)
	_, err = serv.BalanceRepo.RepoFinById(ctx, balanceIn.Id)
	if err != nil {
		return
	}
	balanceRes, err = serv.BalanceRepo.RepoUpdate(ctx, balanceIn)
	if err != nil {
		return
	}
	return
}
func (serv *ServBalanceImpl) ServDelete(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,ServDelete", serv)
	fmt.Println(log)
	balanceRes, err = serv.BalanceRepo.RepoDelete(ctx, balanceId)
	if err != nil {
		return
	}
	return
}
