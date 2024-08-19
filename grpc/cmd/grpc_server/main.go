package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/travacry/auth/grpc/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func (s *server) CreateUser(_ context.Context, req *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	fmt.Print(color.RedString("Create: "))
	fmt.Print(color.GreenString("%+v, pass : %s, cpass : %s\n", req.GetInfo(), req.GetPassword(), req.GetPasswordConfirm()))

	return &user_v1.CreateUserResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) GetUser(_ context.Context, req *user_v1.GetUserRequest) (*user_v1.GetUserResponse, error) {
	fmt.Print(color.RedString("Get: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

	return &user_v1.GetUserResponse{
		User: &user_v1.User{
			Id: req.GetId(),
			Info: &user_v1.UserInfo{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Role:  user_v1.Role_USER,
			},
			CreateAt: timestamppb.New(gofakeit.Date()),
			UpdateAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) UpdateUser(_ context.Context, req *user_v1.UpdateUserRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Update: "))
	fmt.Print(color.GreenString("%v\n", req.GetInfo()))

	return &empty.Empty{}, nil
}

func (s *server) DeleteUser(_ context.Context, req *user_v1.DeleteUserRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Delete: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

	return &empty.Empty{}, nil
}
