package main

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	desc "github.com/travacry/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051"
	userID  = 100001
	pass    = "123"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createUser(ctx, client)
	getUser(ctx, client)
	updateUser(ctx, client)
	deleteUser(ctx, client)
}

func createUser(ctx context.Context, client desc.UserV1Client) {
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

func getUser(ctx context.Context, client desc.UserV1Client) {
	getRequest, err := client.Get(ctx, &desc.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}
	log.Printf(color.RedString("Get user:\n"), color.GreenString("%+v", getRequest.GetUser()))
}

func updateUser(ctx context.Context, client desc.UserV1Client) {
	_, err := client.Update(ctx, &desc.UpdateRequest{
		Info: &desc.UpdateUserInfo{
			Id:    userID,
			Name:  wrapperspb.String(gofakeit.Name()),
			Email: wrapperspb.String(gofakeit.Email()),
			Role:  desc.Role_USER,
		},
	})
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}
	log.Print(color.RedString("Update user.\n"))
}

func deleteUser(ctx context.Context, client desc.UserV1Client) {
	_, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: userID,
	})
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}
	log.Print(color.RedString("Delete user.\n"))
}
