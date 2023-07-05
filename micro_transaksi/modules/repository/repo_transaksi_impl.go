package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"micro_transaksi/modules/model"
)

type RepoTransaksiImpl struct {
	DB *sql.DB
}

func NewRepoTransaksiImpl(db *sql.DB) RepoTransaksi {
	return &RepoTransaksiImpl{
		DB: db,
	}
}
func (repo *RepoTransaksiImpl) RepoCreate(ctx context.Context, transaksiIn model.Transaksi) (err error) {
	logCtx := fmt.Sprintf("%T, RepoCreate", repo)
	log.Println(logCtx)
	sql := "INSERT INTO transaksi(id_product,id_user,qty,total,tgl_transaksi) VALUES($1,$2,$3,$4,$5)"
	_, err = repo.DB.ExecContext(ctx, sql, transaksiIn.Id_Product, transaksiIn.Id_User, transaksiIn.Quantity, transaksiIn.Total, transaksiIn.Tgl_Transaksi)
	if err != nil {
		return
	}
	return
}
