package helper

import (
	"micro_user/modules/model"
	"micro_user/modules/proto"
)

func ProtoUserToStruct(userIn *proto.User) model.User {
	return model.User{
		Name:     userIn.Name,
		Email:    userIn.Email,
		Username: userIn.Username,
		Password: userIn.Password,
		No_hp:    userIn.NoHp,
		Alamat:   userIn.Alamat,
		Jenkel:   userIn.Jenkel,
	}

}
