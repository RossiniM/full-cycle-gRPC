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

	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient) {
	var reqs []*pb.User

	reqs = append(reqs, &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}, &pb.User{
		Id:    "1",
		Name:  "Amanda",
		Email: "amanda@j.com",
	}, &pb.User{
		Id:    "2",
		Name:  "Wesley",
		Email: "wesley@j.com",
	})
	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Could not creating request: %v\n", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}
	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Could not receiving stream : %v\n", err)
	}

	fmt.Println(resp)
}
