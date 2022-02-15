package main

import (
	"github.com/RossiniM/full-cycle-gRPC/pb"
	"github.com/RossiniM/full-cycle-gRPC/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could no connect %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{})
	reflection.Register(grpcServer)
	if err2 := grpcServer.Serve(lis); err2 != nil {
		log.Fatalf("could no server %v", err2)
	}
}
