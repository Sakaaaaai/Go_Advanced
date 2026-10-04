package main

import (
	"context"
	"fmt"
)

// ========= 自動生成されるコードのモック =========
type UserRequest struct { Id int32 }
type User struct { Id int32; Name string }

type grpcClientConn struct{} // grpc.ClientConn のモック

type UserServiceClient interface {
	GetUser(ctx context.Context, in *UserRequest) (*User, error)
}
type userServiceClient struct { cc *grpcClientConn }
func (c *userServiceClient) GetUser(ctx context.Context, in *UserRequest) (*User, error) {
	// サーバーを呼び出すふりをしてモックデータを返す
	return &User{Id: in.Id, Name: "Alice (from mock server)"}, nil
}
func NewUserServiceClient(cc *grpcClientConn) UserServiceClient {
	return &userServiceClient{cc}
}
// ===============================================

func CallGetUser(conn *grpcClientConn) error {
	// 1. クライアントオブジェクトを生成する
	client := ___(conn)

	// 2. サーバーの関数を呼び出す（ローカル関数を呼ぶようにRPC通信が行われる）
	req := &UserRequest{Id: 1}
	res, err := client.___(___.Background(), req)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	// 3. 結果を出力する
	fmt.Printf("User found: %s\n", res.___)
	return nil
}

func main() {
	// 実際のgRPCクライアント接続イメージ（今回は実行しません）
	// conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// defer conn.Close()
	
	conn := &grpcClientConn{}
	err := CallGetUser(conn)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
