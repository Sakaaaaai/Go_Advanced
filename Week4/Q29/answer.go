//go:build ignore

package main

import (
	"context"
	"fmt"
)

type codesType int
const (
	OK codesType = 0
	InvalidArgument codesType = 3
	NotFound codesType = 5
)

func statusError(c codesType, msg string) error {
	return fmt.Errorf("rpc error: code = %v desc = %s", c, msg)
}

type UserRequest struct { Id string }
type User struct { Id string; Name string }

type server struct{}

func (s *server) GetUser(ctx context.Context, req *UserRequest) (*User, error) {
	if req.Id == "" {
		return nil, statusError(InvalidArgument, "ID is required")
	}

	db := map[string]string{"1": "Alice", "2": "Bob"}
	name, exists := db[req.Id]

	if !exists {
		return nil, statusError(NotFound, "User not found")
	}

	return &User{Id: req.Id, Name: name}, nil
}

func main() {
	s := &server{}

	_, err1 := s.GetUser(context.Background(), &UserRequest{Id: ""})
	fmt.Println("Test 1 (Empty ID):", err1)

	_, err2 := s.GetUser(context.Background(), &UserRequest{Id: "99"})
	fmt.Println("Test 2 (Not Found):", err2)

	user, err3 := s.GetUser(context.Background(), &UserRequest{Id: "1"})
	fmt.Printf("Test 3 (Success): user=%+v, err=%v\n", user, err3)
}
