//go:build ignore

package main

import (
	"context"
	"fmt"
)

type UserRequest struct { Id int32 }
type User struct { Id int32; Name string }
type grpcClientConn struct{}

type UserServiceClient interface {
	GetUser(ctx context.Context, in *UserRequest) (*User, error)
}
type userServiceClient struct { cc *grpcClientConn }
func (c *userServiceClient) GetUser(ctx context.Context, in *UserRequest) (*User, error) {
	return &User{Id: in.Id, Name: "Alice (from mock server)"}, nil
}
func NewUserServiceClient(cc *grpcClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func CallGetUser(conn *grpcClientConn) error {
	client := NewUserServiceClient(conn)

	req := &UserRequest{Id: 1}
	res, err := client.GetUser(context.Background(), req)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	fmt.Printf("User found: %s\n", res.Name)
	return nil
}

func main() {
	conn := &grpcClientConn{}
	err := CallGetUser(conn)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
