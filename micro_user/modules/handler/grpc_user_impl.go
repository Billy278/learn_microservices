package handler

import (
	"context"
	"micro_user/modules/proto"
	"micro_user/modules/services"
	"micro_user/pkg/helper"
	"net/http"
)

type GRPCuserImpl struct {
	proto.UnimplementedUsersServer
	UserServ services.UserServ
}

func NewGRPCuserImpl(userserv services.UserServ) *GRPCuserImpl {
	return &GRPCuserImpl{
		UserServ: userserv,
	}
}

func (ctrl *GRPCuserImpl) Register(ctx context.Context, UserReq *proto.User) (*proto.ResponseRegister, error) {
	resIn := helper.ProtoUserToStruct(UserReq)
	Res, err := ctrl.UserServ.ServCreateUser(ctx, resIn)
	if err != nil {
		return &proto.ResponseRegister{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}

	return &proto.ResponseRegister{
		Code: http.StatusOK,
		Data: &proto.User{
			Username: Res.Username,
		},
	}, nil
}
func (ctrl *GRPCuserImpl) Login(ctx context.Context, UserReq *proto.LoginIn) (*proto.ResponseLogin, error) {
	return &proto.ResponseLogin{}, nil
}
