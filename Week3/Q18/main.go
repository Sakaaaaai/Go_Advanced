package main

import (
	"context"
	"fmt"
)

// ========= モック環境 =========
// grpcのパッケージで定義されている型をモックします
type UnaryServerInfo struct {
	FullMethod string
}
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)
// =============================

// インターセプタ関数の実装
func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *UnaryServerInfo,
	handler ___, // 本来の処理を行う関数の型
) (interface{}, error) {

	// 1. ハンドラ実行前の処理（ロギング）
	fmt.Printf("[Global Interceptor] Method called: %s\n", info.___)

	// 2. 本来のハンドラ（RPCメソッド）を呼び出す
	res, err := ___(ctx, req)

	// 3. ハンドラ実行後の処理（ロギング）
	fmt.Println("[Global Interceptor] Method execution finished.")

	// 4. そのまま結果を返す
	return res, err
}

func main() {
	// 動作確認用モック呼び出し
	dummyHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		fmt.Println("  -> Inside actual handler (e.g. GetUser)")
		return "Success", nil
	}
	
	info := &UnaryServerInfo{FullMethod: "/user.UserService/GetUser"}

	// インターセプタ越しにダミーハンドラを呼んでみる
	res, err := loggingInterceptor(context.Background(), "dummy_request", info, dummyHandler)
	fmt.Printf("Result: %v, Error: %v\n", res, err)
}
