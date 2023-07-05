package repository

import (
	"context"
	"micro_product/modules/model"
)

type RepoProduct interface {
	ShowProduct(ctx context.Context) (productRes []model.Product, err error)
	CreateProduct(ctx context.Context, productIn model.Product) (productRes model.Product, err error)
	FindByProduct(ctx context.Context, productId uint64) (productRes model.Product, err error)
	UpdateProduct(ctx context.Context, productIn model.Product) (productRes model.Product, err error)
	DeleteProduct(ctx context.Context, productId uint64) (productRes model.Product, err error)
	UpdateStok(ctx context.Context, productIn model.Product) (productRes model.Product, err error)
}
