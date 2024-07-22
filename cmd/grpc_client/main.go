package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	desc "github.com/travacry/auth/pkg/user_v1"
)

const (
	address = "localhost:50051"
	userID  = 100001
	pass    = "123"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to connect to server: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	client := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// log.Printf("failed to get user by id: %v", err)
	_, err = createUser(ctx, client)

	if err != nil {
		log.Print(err)
	}

	_, err = getUser(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = updateUser(ctx, client)

	if err != nil {
		log.Print(err)
	}

	err = deleteUser(ctx, client)

	if err != nil {
		log.Print(err)
	}
}

func createUser(ctx context.Context, client desc.UserV1Client) (*desc.CreateResponse, error) {

	createResponse, err := client.Create(ctx, &desc.CreateRequest{
		Info: &desc.UserInfo{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Role:  desc.Role_user,
		},
		Password:        pass,
		PasswordConfirm: pass,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	fmt.Print(color.RedString("Create user: "))
	fmt.Printf(color.GreenString("%+d\n", createResponse.GetId()))
	return createResponse, nil
}

func getUser(ctx context.Context, client desc.UserV1Client) (*desc.GetResponse, error) {

	getResponse, err := client.Get(ctx, &desc.GetRequest{Id: userID})

	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}

	fmt.Print(color.RedString("Get user: "))
	fmt.Printf(color.GreenString("%+v\n", getResponse.GetUser()))
	return getResponse, nil
}

func updateUser(ctx context.Context, client desc.UserV1Client) error {
	_, err := client.Update(ctx, &desc.UpdateRequest{
		Info: &desc.UpdateUserInfo{
			Id:    userID,
			Name:  wrapperspb.String(gofakeit.Name()),
			Email: wrapperspb.String(gofakeit.Email()),
			Role:  desc.Role_user,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	fmt.Print(color.RedString("Update user.\n"))
	return nil
}

func deleteUser(ctx context.Context, client desc.UserV1Client) error {
	_, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: userID,
	})

	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	fmt.Print(color.RedString("Delete user.\n"))
	return nil
}
