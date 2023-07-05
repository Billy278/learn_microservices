package services

import (
	"context"
	"micro_email/modules/model"
)

type ServEmail interface {
	ServCreate(ctx context.Context, emailIn model.Email) (emailRes model.Email, err error)
}
