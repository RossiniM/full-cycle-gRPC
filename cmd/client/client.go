package main

import (
	"context"
	"fmt"
	"github.com/RossiniM/full-cycle-gRPC/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	start := time.Now()
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v\n", err)
	}
	defer connection.Close()
	client := pb.NewUserServiceClient(connection)

	addUserVerbose(client)
	//for i := 0; i < 10000; i++ {
	//	AddUser(client)
	//	fmt.Println(i)
	//}
	elapsed := time.Since(start)
	log.Printf("Time Elapsed %s", elapsed)

}

// AddUser unary communication
func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}
	_, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v\n", err)
	}
}

// AddUserVerbose client for stream server
func addUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}
	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v\n", err)
	}
	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("could not received")
		}
		fmt.Println("Status:", stream.Status)
	}
}
