package server

import (
	"log"
	"micro_product/modules/proto"
	"micro_product/modules/router"
	"micro_product/pkg/middleware"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuthZ())

	ctrl := NewSetup()
	router.NewRouter(s, ctrl.CtrlProduct)
	s.Run(":9090")

}

func NewGRPCServer() {

	srv := grpc.NewServer(middleware.WithMiddlewareUnarryInceptor())
	ctrl := GRPCSetup()
	proto.RegisterProductsServer(srv, ctrl)
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("server failed to run at port 9090 ")
	}
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
