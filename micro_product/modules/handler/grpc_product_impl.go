package handler

import (
	"context"
	"micro_product/modules/proto"
	"micro_product/modules/services"
	"micro_product/pkg/helper"
	"net/http"
)

type HandlerProductImpl struct {
	ProductServ services.SrvProduct
	proto.UnimplementedProductsServer
}

func NewHandlerProductImpl(productserv services.SrvProduct) *HandlerProductImpl {
	return &HandlerProductImpl{
		ProductServ: productserv,
	}
}

func (ctrl *HandlerProductImpl) Show(ctx context.Context, reqIn *proto.Product) (*proto.Response, error) {
	res, err := ctrl.ProductServ.SrvShow(ctx)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusAccepted,
		Data: helper.StructToprotoProducts(res),
	}, nil
}
func (ctrl *HandlerProductImpl) Create(ctx context.Context, reqIn *proto.Product) (*proto.Response, error) {
	res, err := ctrl.ProductServ.SrvCreate(ctx, helper.ProtoProductToStruct(reqIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusCreated,
		Data: &proto.ProductList{
			List: []*proto.Product{
				helper.StructToprotoProduct(res),
			},
		},
	}, nil
}
func (ctrl *HandlerProductImpl) FindByid(ctx context.Context, reqIn *proto.Product) (*proto.Response, error) {
	res, err := ctrl.ProductServ.SrvFindBy(ctx, reqIn.Id)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusAccepted,
		Data: &proto.ProductList{
			List: []*proto.Product{
				helper.StructToprotoProduct(res),
			},
		},
	}, nil
}
func (ctrl *HandlerProductImpl) Update(ctx context.Context, reqIn *proto.Product) (*proto.Response, error) {
	res, err := ctrl.ProductServ.SrvUpdate(ctx, helper.ProtoProductToStruct(reqIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.ProductList{
			List: []*proto.Product{
				helper.StructToprotoProduct(res),
			},
		},
	}, nil
}
func (ctrl *HandlerProductImpl) Delete(ctx context.Context, reqIn *proto.Product) (*proto.Response, error) {
	res, err := ctrl.ProductServ.SrvDelete(ctx, reqIn.Id)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.ProductList{
			List: []*proto.Product{
				helper.StructToprotoProduct(res),
			},
		},
	}, nil
}
func (ctrl *HandlerProductImpl) UpdateStock(ctx context.Context, reqIn *proto.Product) (*proto.Response, error) {
	res, err := ctrl.ProductServ.SrvUpdateStock(ctx, helper.ProtoProductToStruct(reqIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.ProductList{
			List: []*proto.Product{
				helper.StructToprotoProduct(res),
			},
		},
	}, nil
}
