package model

type User struct {
	Id       int
	Name     string `json:"name" validate:"required"`
	Email    string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	No_hp    string `json:"no_hp" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	Jenkel   string `json:"jenkel" validate:"required"`
}
