package services

import (
	"context"
	"fmt"
	"github.com/RossiniM/full-cycle-gRPC/pb"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

//AddUser server using unary communication
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	fmt.Printf("id:%s, name:%s, email:%s", user.Id, user.Name, user.Email)
	return &user, nil

}

//AddUserVerbose server using stream
func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting at database",
		User: &pb.User{
			Id:    req.Id,
			Name:  req.Name,
			Email: req.Email,
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    req.Id,
			Name:  req.Name,
			Email: req.Email,
		},
	})
	time.Sleep(time.Second * 3)
	return nil

}
