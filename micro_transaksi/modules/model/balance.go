package model

type Balance struct {
	Id     uint64  `json:"id"`
	UserId uint64  `json:"user_id"`
	Saldo  float64 `json:"saldo"`
}
