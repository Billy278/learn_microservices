package main

import (
	"context"
	"fmt"
	"micro_balance/modules/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func serviceClient() proto.BalancesClient {
	port := ":8090"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect to", port, err)
	}
	return proto.NewBalancesClient(conn)

}
func main() {
	balance := serviceClient()

	//md:= metadata.Pairs("z", "abc")
	mz := metadata.MD{
		"key": {"abc"},
	}
	c := metadata.NewOutgoingContext(context.Background(), mz)
	res, err := balance.Show(c, &proto.Balance{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
