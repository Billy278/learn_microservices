package server

import (
	"log"
	"micro_balance/modules/proto"
	"micro_balance/modules/router"
	"micro_balance/pkg/middleware"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuth())
	ctrl := NewSetup()
	router.NewRouterBlnc(s, ctrl.CtrlBalance)
	s.Run(":8090")
}

func NewGRPCServer() {

	srv := grpc.NewServer()
	ctrl := GRPCSetup()
	proto.RegisterBalancesServer(srv, ctrl)
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("server failed to run at port 8090 ")
	}
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
