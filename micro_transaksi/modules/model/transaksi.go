package model

import "time"

type Transaksi struct {
	Id            uint64     `json:"id"`
	Id_Product    uint64     `json:"id_product" validate:"required,numeric"`
	Id_User       uint64     `json:"id_user"`
	Quantity      uint64     `json:"qty" validate:"required,numeric"`
	Total         float64    `json:"total"`
	Tgl_Transaksi *time.Time `json:"tgl_transaksi"`
}
