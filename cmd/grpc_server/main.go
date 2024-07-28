package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/travacry/auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func (s *server) CreateUser(_ context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	fmt.Print(color.RedString("Create: "))
	fmt.Print(color.GreenString("%+v, pass : %s, cpass : %s\n", req.GetInfo(), req.GetPassword(), req.GetPasswordConfirm()))

	return &desc.CreateUserResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) GetUser(_ context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	fmt.Print(color.RedString("Get: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

	return &desc.GetUserResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Role:  desc.Role_USER,
			},
			CreateAt: timestamppb.New(gofakeit.Date()),
			UpdateAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) UpdateUser(_ context.Context, req *desc.UpdateUserRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Update: "))
	fmt.Print(color.GreenString("%v\n", req.GetInfo()))

	return &empty.Empty{}, nil
}

func (s *server) DeleteUser(_ context.Context, req *desc.DeleteUserRequest) (*empty.Empty, error) {
	fmt.Print(color.RedString("Delete: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

	return &empty.Empty{}, nil
}
