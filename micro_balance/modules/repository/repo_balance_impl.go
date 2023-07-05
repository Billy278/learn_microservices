package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"micro_balance/modules/model"
)

type RepoBalanceImpl struct {
	DB *sql.DB
}

func NewRepoBalanceImpl(db *sql.DB) RepoBalance {
	return &RepoBalanceImpl{
		DB: db,
	}
}

func (repo *RepoBalanceImpl) RepoShow(ctx context.Context) (balanceRes []model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoShow", repo)
	fmt.Println(log)
	sql := "SELECT id,user_id,saldo FROM balance"
	row, err := repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	balance := model.Balance{}
	for row.Next() {
		err = row.Scan(&balance.Id, &balance.UserId, &balance.Saldo)
		if err != nil {
			return
		}
		balanceRes = append(balanceRes, balance)
	}
	return
}
func (repo *RepoBalanceImpl) RepoCreate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoCreate", repo)
	fmt.Println(log)
	sql := "INSERT INTO balance(user_id,saldo) VALUES ($1,$2)"
	_, err = repo.DB.ExecContext(ctx, sql, balanceIn.UserId, balanceIn.Saldo)
	if err != nil {
		return
	}
	balanceRes.UserId = balanceIn.UserId
	return
}
func (repo *RepoBalanceImpl) RepoFinById(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoFinById", repo)
	fmt.Println(log)
	fmt.Println(balanceId)
	sql := "SELECT id,user_id,saldo FROM balance WHERE id=$1"
	row, err := repo.DB.QueryContext(ctx, sql, balanceId)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&balanceRes.Id, &balanceRes.UserId, &balanceRes.Saldo)
		if err != nil {
			return
		}
		return
	} else {
		err = errors.New("ID NOT FOUND")
		return
	}
}
func (repo *RepoBalanceImpl) RepoUpdate(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoUpdate", repo)
	fmt.Println(log)
	sql := "UPDATE balance SET saldo=$1 WHERE id=$2"
	_, err = repo.DB.ExecContext(ctx, sql, balanceIn.Saldo, balanceIn.Id)
	if err != nil {
		return
	}
	balanceRes.Id = balanceIn.Id
	return
}
func (repo *RepoBalanceImpl) RepoUpdateByUser(ctx context.Context, balanceIn model.Balance) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoUpdateByUser", repo)
	fmt.Println(log)
	fmt.Println(balanceIn.Saldo)
	sql := "UPDATE balance SET saldo=$1 WHERE user_id=$2"
	_, err = repo.DB.ExecContext(ctx, sql, balanceIn.Saldo, balanceIn.UserId)
	if err != nil {
		return
	}
	balanceRes.Id = balanceIn.Id
	return
}
func (repo *RepoBalanceImpl) RepoDelete(ctx context.Context, balanceId uint64) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoDelete", repo)
	fmt.Println(log)
	sql := "DELETE FROM balance WHERE id=$1"
	_, err = repo.DB.ExecContext(ctx, sql, balanceId)
	if err != nil {
		return
	}
	balanceRes.Id = balanceId
	return
}
func (repo *RepoBalanceImpl) RepoFinByIdUser(ctx context.Context, userId uint64) (balanceRes model.Balance, err error) {
	log := fmt.Sprintf("%T,RepoFinByIdUser", repo)
	fmt.Println(log)
	sql := "SELECT id,user_id,saldo FROM balance WHERE user_id=$1"
	row, err := repo.DB.QueryContext(ctx, sql, userId)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&balanceRes.Id, &balanceRes.UserId, &balanceRes.Saldo)
		if err != nil {
			return
		}
		return
	} else {
		err = errors.New("ID NOT FOUND")
		return
	}
}
