package helper

import (
	"micro_product/modules/model"
	"micro_product/modules/proto"
)

func ProtoProductToStruct(productIn *proto.Product) model.Product {
	return model.Product{
		Id:     productIn.Id,
		UserId: productIn.Userid,
		Name:   productIn.Name,
		Stock:  productIn.Stock,
		Harga:  float64(productIn.Harga),
	}

}

func StructToprotoProduct(productIn model.Product) *proto.Product {
	return &proto.Product{
		Id:     productIn.Id,
		Userid: productIn.UserId,
		Name:   productIn.Name,
		Stock:  productIn.Stock,
		Harga:  float32(productIn.Harga),
	}

}

func StructToprotoProducts(productIn []model.Product) *proto.ProductList {
	protoProducts := &proto.ProductList{}
	for _, v := range productIn {
		protoProducts.List = append(protoProducts.List, StructToprotoProduct(v))
	}
	return protoProducts

}
