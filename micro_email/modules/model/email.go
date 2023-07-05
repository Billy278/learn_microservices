package model

type Email struct {
	Id             uint64
	Sender         string
	Name_Receiver  string  `json:"name_receiver" validate:"required"`
	Name_product   string  `json:"name_product" validate:"required"`
	Harga          float64 `json:"harga" validate:"required"`
	Qty            uint64  `json:"qty" validate:"required"`
	Total          float64 `json:"total" validate:"required"`
	Email_Receiver string  `json:"email_receiver" validate:"required"`
}
