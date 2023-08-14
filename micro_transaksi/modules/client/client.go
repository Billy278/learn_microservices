package client

import (
	"fmt"
	"micro_transaksi/modules/proto"

	"google.golang.org/grpc"
)

func ServiceClientBalance() proto.BalancesClient {
	port := ":8090"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect to", port, err)
	}

	return proto.NewBalancesClient(conn)
}

func ServiceClientProduct() proto.ProductsClient {
	port := ":9090"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect to", port, err)
	}

	return proto.NewProductsClient(conn)
}
