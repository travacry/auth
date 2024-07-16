package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	desc "github.com/travacry/auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	log.Printf(color.RedString("Create User:\n"),
		color.GreenString("info : %+v, pass : %s, cpass : %s", req.GetInfo(), req.GetPassword(), req.GetPasswordConfirm()))

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	log.Printf(color.RedString("Get User:\n"),
		color.GreenString("info : %d", req.GetId()))

	return &desc.GetResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Role:  desc.Role_user,
			},
			CreateAt: timestamppb.New(gofakeit.Date()),
			UpdateAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {

	log.Printf(color.RedString("Update User:\n"),
		color.GreenString("info : %+v", req.GetInfo()))

	return &empty.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	log.Printf(color.RedString("Delete User:\n"),
		color.GreenString("info : %+v", req.GetId()))

	return &empty.Empty{}, nil
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
		log.Panicf("failed to serve: %v", err)
	}
}
