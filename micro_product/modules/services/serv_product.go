package services

import (
	"context"
	"micro_product/modules/model"
)

type SrvProduct interface {
	SrvShow(ctx context.Context) (productRes []model.Product, err error)
	SrvCreate(ctx context.Context, productIn model.Product) (productRes model.Product, err error)
	SrvFindBy(ctx context.Context, productId uint64) (productRes model.Product, err error)
	SrvUpdate(ctx context.Context, productIn model.Product) (productRes model.Product, err error)
	SrvDelete(ctx context.Context, productId uint64) (productRes model.Product, err error)
	SrvUpdateStock(ctx context.Context, productIn model.Product) (productRes model.Product, err error)
}
