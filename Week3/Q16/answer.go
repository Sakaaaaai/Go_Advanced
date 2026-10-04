//go:build ignore

package main

import (
	"context"
	"fmt"
)

type UserRequest struct { Id int32 }
type User struct {
	Id   int32
	Name string
}

type UserServiceServer interface {
	GetUser(context.Context, *UserRequest) (*User, error)
}

type server struct{}

func (s *server) GetUser(ctx context.Context, req *UserRequest) (*User, error) {
	if req.Id == 1 {
		return &User{Id: 1, Name: "Alice"}, nil
	}
	
	return &User{Id: req.Id, Name: "Unknown"}, nil
}

func main() {
	fmt.Println("This is a package for gRPC Server logic.")
}
