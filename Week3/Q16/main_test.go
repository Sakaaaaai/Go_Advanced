package main

import (
	"context"
	"testing"
)

func TestGRPCServer(t *testing.T) {
	s := &server{}

	// Interface check
	var _ UserServiceServer = s

	user, err := s.GetUser(context.Background(), &UserRequest{Id: 1})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if user.Name != "Alice" {
		t.Errorf("Expected user name to be 'Alice', got '%s'", user.Name)
	}

	user2, err := s.GetUser(context.Background(), &UserRequest{Id: 99})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if user2.Name != "Unknown" {
		t.Errorf("Expected user name to be 'Unknown', got '%s'", user2.Name)
	}
}
