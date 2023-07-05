package repository

import (
	"context"
	"database/sql"
	"fmt"
	"micro_email/modules/model"
)

type RepoEmailImpl struct {
	DB *sql.DB
}

func NewRepoEmailImpl(db *sql.DB) RepoEmail {
	return &RepoEmailImpl{
		DB: db,
	}
}
func (repo *RepoEmailImpl) RepoCreate(ctx context.Context, emailIn model.Email) (emailRes model.Email, err error) {
	log := fmt.Sprintf("%T,RepoCreate", repo)
	fmt.Println(log)
	emailIn.Sender = "tokoserba@gmail.com"
	sql := "INSERT INTO email(sender,name_receiver,name_product,harga,qty,total,email_receiver) VALUES($1,$2,$3,$4,$5,$6,$7)"
	_, err = repo.DB.ExecContext(ctx, sql, emailIn.Sender, emailIn.Name_Receiver, emailIn.Name_product, emailIn.Harga, emailIn.Qty, emailIn.Total, emailIn.Email_Receiver)
	if err != nil {
		return
	}
	emailRes.Name_Receiver = emailIn.Name_Receiver
	return
}
