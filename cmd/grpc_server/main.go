package main

import (
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/travacry/auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	fmt.Print(color.RedString("Create: "))
	fmt.Print(color.GreenString("%+v, pass : %s, cpass : %s\n", req.GetInfo(), req.GetPassword(), req.GetPasswordConfirm()))

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	fmt.Print(color.RedString("Get: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

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

	fmt.Print(color.RedString("Update: "))
	fmt.Print(color.GreenString("%v\n", req.GetInfo()))

	return &empty.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	fmt.Print(color.RedString("Delete: "))
	fmt.Print(color.GreenString("%d\n", req.GetId()))

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
		log.Printf("failed to serve: %v", err)
	}
}
