package server

import (
	"log"
	"net"

	userpb "github.com/ankitanwar/assignment/proto"
	services "github.com/ankitanwar/assignment/server/services"
	"google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatalln("Unable To start the server")
		return
	}
	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	userpb.RegisterUserServicesServer(srv, services.UserService)
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalln("Unable To Listen")
		return
	}

}
