package helper

import (
	"micro_balance/modules/model"
	"micro_balance/modules/proto"
)

func ProtoBalanceToStruct(balanceIn *proto.Balance) model.Balance {
	return model.Balance{
		Id:     balanceIn.Id,
		UserId: balanceIn.Userid,
		Saldo:  float64(balanceIn.Saldo),
	}

}

func StructToprotoBalance(balanceIn model.Balance) *proto.Balance {
	return &proto.Balance{
		Id:     balanceIn.Id,
		Userid: balanceIn.UserId,
		Saldo:  float32(balanceIn.Saldo),
	}

}

func StructToprotoBalances(balanceIn []model.Balance) *proto.BalanceList {
	protoBalances := &proto.BalanceList{}
	for _, v := range balanceIn {
		protoBalances.List = append(protoBalances.List, StructToprotoBalance(v))
	}
	return protoBalances

}
