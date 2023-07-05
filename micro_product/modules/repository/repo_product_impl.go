package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"micro_product/modules/model"
)

type RepoProductImpl struct {
	DB *sql.DB
}

func NewRepoProductImpl(db *sql.DB) RepoProduct {
	return &RepoProductImpl{
		DB: db,
	}
}
func (r_product *RepoProductImpl) ShowProduct(ctx context.Context) (productRes []model.Product, err error) {
	log := fmt.Sprintf("%T,ShowProduct", r_product)
	fmt.Println(log)
	sql := "SELECT id,user_id,name,harga,stock FROM product"
	row, err := r_product.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	product := model.Product{}
	for row.Next() {
		err = row.Scan(&product.Id, &product.UserId, &product.Name, &product.Harga, &product.Stock)
		if err != nil {
			return
		}
		productRes = append(productRes, product)
	}
	return
}
func (r_product *RepoProductImpl) CreateProduct(ctx context.Context, productIn model.Product) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,CreateProduct", r_product)
	fmt.Println(log)
	sql := "INSERT INTO product(user_id,name,harga,stock) VALUES ($1,$2,$3,$4)"
	_, err = r_product.DB.ExecContext(ctx, sql, productIn.UserId, productIn.Name, productIn.Harga, productIn.Stock)
	if err != nil {
		return
	}
	productRes.Name = productIn.Name
	return
}
func (r_product *RepoProductImpl) FindByProduct(ctx context.Context, productId uint64) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,FindByProduct", r_product)
	fmt.Println(log)

	sql := "SELECT id,user_id,name,harga,stock FROM product WHERE id=$1"
	row, err := r_product.DB.QueryContext(ctx, sql, productId)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&productRes.Id, &productRes.UserId, &productRes.Name, &productRes.Harga, &productRes.Stock)
		if err != nil {
			return
		}
		return
	} else {
		err = errors.New("PRODUCT NOT FOUND")
		return
	}
}
func (r_product *RepoProductImpl) UpdateProduct(ctx context.Context, productIn model.Product) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,UpdateProduct", r_product)
	fmt.Println(log)
	sql := "UPDATE product SET name=$1,harga=$2,stock=$3 WHERE id=$4"
	_, err = r_product.DB.ExecContext(ctx, sql, productIn.Name, productIn.Harga, productIn.Stock, productIn.Id)
	if err != nil {
		return
	}
	productRes.Id = productIn.Id
	return
}

func (r_product *RepoProductImpl) UpdateStok(ctx context.Context, productIn model.Product) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,UpdateStok", r_product)
	fmt.Println(log)
	sql := "UPDATE product SET stock=$1 WHERE id=$2"
	_, err = r_product.DB.ExecContext(ctx, sql, productIn.Stock, productIn.Id)
	if err != nil {
		return
	}
	productRes.Id = productIn.Id
	return
}
func (r_product *RepoProductImpl) DeleteProduct(ctx context.Context, productId uint64) (productRes model.Product, err error) {
	log := fmt.Sprintf("%T,DeleteProduct", r_product)
	fmt.Println(log)
	sql := "DELETE FROM product WHERE id=$1"
	_, err = r_product.DB.ExecContext(ctx, sql, productId)
	if err != nil {
		return
	}
	productRes.Id = productId
	return
}
