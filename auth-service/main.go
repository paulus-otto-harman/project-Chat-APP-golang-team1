package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"project/auth-service/infra"
	pb "project/auth-service/proto"
	"project/auth-service/service"
)

func main() {
	var err error
	var ctx *infra.ServiceContext
	ctx, err = infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}
	log.Println(ctx)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, &service.AuthService{})

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
