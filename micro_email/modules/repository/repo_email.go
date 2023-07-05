package repository

import (
	"context"
	"micro_email/modules/model"
)

type RepoEmail interface {
	RepoCreate(ctx context.Context, emailIn model.Email) (emailRes model.Email, err error)
}
