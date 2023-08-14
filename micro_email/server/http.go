package server

import (
	"log"
	"micro_email/modules/proto"
	"micro_email/modules/router"
	"micro_email/pkg/middleware"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewServer() {
	s := gin.Default()
	s.Use(gin.Recovery(), gin.Logger(), middleware.BearerOAuthZ())
	ctrl := NewSetup()
	router.NewRouterEmail(s, ctrl.CtrlEmail)
	s.Run(":6060")
}

func NewGRPCServer() {

	srv := grpc.NewServer(middleware.WithMiddlewareUnarryInceptor())
	ctrl := GRPCSetup()
	proto.RegisterEmailSrvServer(srv, ctrl)
	l, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatal("server failed to run at port 6060 ")
	}
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
