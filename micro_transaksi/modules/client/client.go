package client

import (
	"fmt"
	"micro_transaksi/modules/proto"
	"os"

	"google.golang.org/grpc"
)

func ServiceClientBalance() proto.BalancesClient {
	port := fmt.Sprintf("%v:%v", os.Getenv("hostBalance"), os.Getenv("PortBalance"))
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect to", port, err)
	}

	return proto.NewBalancesClient(conn)
}

func ServiceClientProduct() proto.ProductsClient {
	port := fmt.Sprintf("%v:%v", os.Getenv("hostProduct"), os.Getenv("PortProduct"))
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect to", port, err)
	}

	return proto.NewProductsClient(conn)
}
