//go:build ignore

package main

import (
	"context"
	"fmt"
)

type UnaryServerInfo struct {
	FullMethod string
}
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *UnaryServerInfo,
	handler UnaryHandler,
) (interface{}, error) {

	fmt.Printf("[Global Interceptor] Method called: %s\n", info.FullMethod)

	res, err := handler(ctx, req)

	fmt.Println("[Global Interceptor] Method execution finished.")

	return res, err
}

func main() {
	dummyHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		fmt.Println("  -> Inside actual handler (e.g. GetUser)")
		return "Success", nil
	}
	
	info := &UnaryServerInfo{FullMethod: "/user.UserService/GetUser"}

	res, err := loggingInterceptor(context.Background(), "dummy_request", info, dummyHandler)
	fmt.Printf("Result: %v, Error: %v\n", res, err)
}
