package server

import (
	"log"
	"micro_user/modules/proto"
	routerUser "micro_user/modules/router/v1/user"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewServer() {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())
	ctrl := NewSetup()
	routerUser.NewUserRouter(r, ctrl.UserCtrl)
	r.Run(":8080")
}

func NewGRPCServer() {

	srv := grpc.NewServer()
	ctrl := GRPCSetup()
	proto.RegisterUsersServer(srv, ctrl)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("server failed to run at port 8080 ")
	}
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
