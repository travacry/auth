package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/travacry/auth/grpc/pkg/user_v1"
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
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to connect to server: %v", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	client := user_v1.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

func createUser(ctx context.Context, client user_v1.UserV1Client) (*user_v1.CreateUserResponse, error) {
	createResponse, err := client.CreateUser(ctx, &user_v1.CreateUserRequest{
		Info: &user_v1.UserInfo{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Role:  user_v1.Role_USER,
		},
		Password:        pass,
		PasswordConfirm: pass,
	})
	if err != nil {
		return nil, createUserError(err)
	}

	fmt.Print(color.RedString("Create user: "))
	fmt.Print(color.GreenString("%+d\n", createResponse.GetId()))
	return createResponse, nil
}
func createUserError(err error) error {
	return fmt.Errorf("failed to create user: %v", err)
}

func getUser(ctx context.Context, client user_v1.UserV1Client) (*user_v1.GetUserResponse, error) {
	getResponse, err := client.GetUser(ctx, &user_v1.GetUserRequest{Id: userID})
	if err != nil {
		return nil, getUserError(err)
	}

	fmt.Print(color.RedString("Get user: "))
	fmt.Print(color.GreenString("%+v\n", getResponse.GetUser()))
	return getResponse, nil
}
func getUserError(err error) error {
	return fmt.Errorf("failed to get user by id: %v", err)
}

func updateUser(ctx context.Context, client user_v1.UserV1Client) error {
	_, err := client.UpdateUser(ctx, &user_v1.UpdateUserRequest{
		Info: &user_v1.UpdateUserInfo{
			Id:   userID,
			Name: wrapperspb.String(gofakeit.Name()),
			Role: user_v1.Role_USER,
		},
	})
	if err != nil {
		return updateUserError(err)
	}

	fmt.Print(color.RedString("Update user.\n"))
	return nil
}
func updateUserError(err error) error {
	return fmt.Errorf("failed to update user: %v", err)
}

func deleteUser(ctx context.Context, client user_v1.UserV1Client) error {
	_, err := client.DeleteUser(ctx, &user_v1.DeleteUserRequest{
		Id: userID,
	})
	if err != nil {
		return deleteUserError(err)
	}

	fmt.Print(color.RedString("Delete user.\n"))
	return nil
}
func deleteUserError(err error) error {
	return fmt.Errorf("failed to delete user: %v", err)
}
