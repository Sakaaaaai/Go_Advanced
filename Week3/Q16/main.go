package main

import (
	"context"
	"fmt"
)

// ========= 自動生成されるコードのモック =========
type UserRequest struct { Id int32 }
type User struct {
	Id   int32
	Name string
}
// このインターフェースを実装する必要があります
type UserServiceServer interface {
	GetUser(context.Context, *UserRequest) (*User, error)
}
// ===============================================

// サーバー実装用の構造体
type server struct{}

// 1. GetUser メソッドを実装する
func (s *server) ___(ctx context.___, req *___) (*___, error) {
	// 2. req.Id に応じて返すUserを変える
	if req.Id == 1 {
		return &User{Id: 1, Name: "Alice"}, nil
	}
	
	return &User{Id: req.Id, Name: "Unknown"}, ___
}

func main() {
	// 実際のgRPCサーバー起動のイメージ（今回は実行しません）
	// lis, _ := net.Listen("tcp", ":50051")
	// s := grpc.NewServer()
	// pb.RegisterUserServiceServer(s, &server{})
	// s.Serve(lis)
	fmt.Println("This is a package for gRPC Server logic.")
}
