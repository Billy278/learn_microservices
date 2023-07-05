package model

type Product struct {
	Id     uint64  `json:"id"`
	UserId uint64  `json:"user_id"`
	Name   string  `json:"name" validate:"required"`
	Stock  uint64  `json:"stock" validate:"required"`
	Harga  float64 `json:"harga" validate:"required"`
}
