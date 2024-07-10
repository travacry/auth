package main

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	desc "github.com/travacry/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	userId  = 100001
	pass    = "123"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	CreateUser(ctx, client)
	GetUser(ctx, client)
	UpdateUser(ctx, client)
	DeleteUser(ctx, client)
}

func CreateUser(ctx context.Context, client desc.UserV1Client) {
	createRequest, err := client.Create(ctx, &desc.CreateRequest{
		Info: &desc.UserInfo{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Role:  desc.Role_USER,
		},
		Password:        pass,
		PasswordConfirm: pass,
	})

	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}
	log.Printf(color.RedString("Create user:\n"), color.GreenString("%+d", createRequest.GetId()))
}

func GetUser(ctx context.Context, client desc.UserV1Client) {
	getRequest, err := client.Get(ctx, &desc.GetRequest{Id: userId})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}
	log.Printf(color.RedString("Get user:\n"), color.GreenString("%+v", getRequest.GetUser()))
}

func UpdateUser(ctx context.Context, client desc.UserV1Client) {
	_, err := client.Update(ctx, &desc.UpdateRequest{
		Info: &desc.UpdateUserInfo{
			Id:    userId,
			Name:  wrapperspb.String(gofakeit.Name()),
			Email: wrapperspb.String(gofakeit.Email()),
			Role:  desc.Role_USER,
		},
	})
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}
	log.Printf(color.RedString("Update user.\n"))
}

func DeleteUser(ctx context.Context, client desc.UserV1Client) {
	_, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: userId,
	})
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}
	log.Printf(color.RedString("Delete user.\n"))
}
