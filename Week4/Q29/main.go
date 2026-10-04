package main

import (
	"context"
	"fmt"

	// gRPCのエラー処理に使うパッケージ（モックとして定義済み）
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

// ========= モック環境 =========
type codesType int
const (
	OK codesType = 0
	InvalidArgument codesType = 3
	NotFound codesType = 5
)

// gRPCのstatus.Errorを模倣
func statusError(c codesType, msg string) error {
	return fmt.Errorf("rpc error: code = %v desc = %s", c, msg)
}

// 自動生成コードのモック
type UserRequest struct { Id string }
type User struct { Id string; Name string }
// =============================

type server struct{}

func (s *server) GetUser(ctx context.Context, req *UserRequest) (*User, error) {
	// 1. バリデーションエラー（IDが空の場合）
	if req.Id == "" {
		// codes.InvalidArgument を返すようにする
		return nil, ___(InvalidArgument, "ID is required")
	}

	// モックデータベースの検索処理
	db := map[string]string{"1": "Alice", "2": "Bob"}
	name, exists := db[req.Id]

	// 2. 見つからなかった場合のエラー
	if !exists {
		// codes.NotFound を返すようにする
		return nil, ___(NotFound, "User not found")
	}

	// 正常に取得できた場合
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
