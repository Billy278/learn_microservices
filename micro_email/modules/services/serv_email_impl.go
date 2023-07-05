package services

import (
	"context"
	"fmt"
	"micro_email/modules/model"
	"micro_email/modules/repository"
)

type ServEmailImpl struct {
	RepoEmail repository.RepoEmail
}

func NewServEmailImpl(repoemail repository.RepoEmail) ServEmail {
	return &ServEmailImpl{
		RepoEmail: repoemail,
	}
}
func (serv *ServEmailImpl) ServCreate(ctx context.Context, emailIn model.Email) (emailRes model.Email, err error) {
	log := fmt.Sprintf("%T,ServCreate", serv)
	fmt.Println(log)
	emailRes, err = serv.RepoEmail.RepoCreate(ctx, emailIn)
	if err != nil {
		return
	}
	return
}
