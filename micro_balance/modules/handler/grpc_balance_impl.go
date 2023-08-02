package handler

import (
	"context"
	"micro_balance/modules/proto"
	"micro_balance/modules/services"
	"micro_balance/pkg/helper"
	"net/http"
)

type GRPCBalanceImpl struct {
	proto.UnimplementedBalancesServer
	BalanceServ services.ServBalance
}

func NewGRPCBalanceImpl(balanceserv services.ServBalance) *GRPCBalanceImpl {
	return &GRPCBalanceImpl{
		BalanceServ: balanceserv,
	}

}

func (ctrl *GRPCBalanceImpl) Show(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {
	res, err := ctrl.BalanceServ.ServShow(ctx)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusAccepted,
		Data: helper.StructToprotoBalances(res),
	}, nil
}
func (ctrl *GRPCBalanceImpl) Create(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {

	res, err := ctrl.BalanceServ.ServCreate(ctx, helper.ProtoBalanceToStruct(balanceIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.BalanceList{
			List: []*proto.Balance{
				&proto.Balance{
					Userid: res.UserId,
				},
			},
		},
	}, nil
}
func (ctrl *GRPCBalanceImpl) FindByid(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {
	res, err := ctrl.BalanceServ.ServFinById(ctx, balanceIn.Id)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusAccepted,
		Data: &proto.BalanceList{
			List: []*proto.Balance{
				helper.StructToprotoBalance(res),
			},
		},
	}, nil
}
func (ctrl *GRPCBalanceImpl) FindByidUser(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {
	res, err := ctrl.BalanceServ.ServFinByIdUser(ctx, balanceIn.Userid)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusAccepted,
		Data: &proto.BalanceList{
			List: []*proto.Balance{
				helper.StructToprotoBalance(res),
			},
		},
	}, nil
}
func (ctrl *GRPCBalanceImpl) Update(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {
	res, err := ctrl.BalanceServ.ServUpdate(ctx, helper.ProtoBalanceToStruct(balanceIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.BalanceList{
			List: []*proto.Balance{
				helper.StructToprotoBalance(res),
			},
		},
	}, nil

}
func (ctrl *GRPCBalanceImpl) UpdateByServer(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {
	res, err := ctrl.BalanceServ.ServUpdateByIdUser(ctx, helper.ProtoBalanceToStruct(balanceIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.BalanceList{
			List: []*proto.Balance{
				helper.StructToprotoBalance(res),
			},
		},
	}, nil
}
func (ctrl *GRPCBalanceImpl) Delete(ctx context.Context, balanceIn *proto.Balance) (*proto.Response, error) {
	res, err := ctrl.BalanceServ.ServDelete(ctx, balanceIn.Id)
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.BalanceList{
			List: []*proto.Balance{
				helper.StructToprotoBalance(res),
			},
		},
	}, nil
}
