package services

import (
	"context"
	"fmt"
	"micro_product/modules/model"
	"micro_product/modules/repository"
)

type SrvProductImpl struct {
	ProductRepo repository.RepoProduct
}

func NewSrvProductImpl(productrepo repository.RepoProduct) SrvProduct {
	return &SrvProductImpl{
		ProductRepo: productrepo,
	}
}

func (srv *SrvProductImpl) SrvShow(ctx context.Context) (productRes []model.Product, err error) {
	log := fmt.Sprintf("%T,SrvShow", srv)
	fmt.Println(log)
	productRes, err = srv.ProductRepo.ShowProduct(ctx)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvCreate(ctx context.Context, productIn model.Product) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,SrvCreate", srv)
	fmt.Println(log)
	productRes, err = srv.ProductRepo.CreateProduct(ctx, productIn)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvFindBy(ctx context.Context, productId uint64) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,SrvFindBy", srv)
	fmt.Println(log)
	productRes, err = srv.ProductRepo.FindByProduct(ctx, productId)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvUpdate(ctx context.Context, productIn model.Product) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,SrvUpdate", srv)
	fmt.Println(log)
	_, err = srv.ProductRepo.FindByProduct(ctx, productIn.Id)
	if err != nil {
		return
	}
	productRes, err = srv.ProductRepo.UpdateProduct(ctx, productIn)
	if err != nil {
		return
	}
	return
}
func (srv *SrvProductImpl) SrvDelete(ctx context.Context, productId uint64) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,SrvDelete", srv)
	fmt.Println(log)
	_, err = srv.ProductRepo.FindByProduct(ctx, productId)
	if err != nil {
		return
	}
	productRes, err = srv.ProductRepo.DeleteProduct(ctx, productId)
	if err != nil {
		return
	}
	return
}

func (srv *SrvProductImpl) SrvUpdateStock(ctx context.Context, productIn model.Product) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,SrvUpdateStock", srv)
	fmt.Println(log)
	_, err = srv.ProductRepo.FindByProduct(ctx, productIn.Id)
	if err != nil {
		return
	}
	productRes, err = srv.ProductRepo.UpdateStok(ctx, productIn)
	if err != nil {
		return
	}
	return
}
