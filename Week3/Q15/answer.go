//go:build ignore

package main

import "fmt"

const ProtoFileContent = `
syntax = "proto3";

message UserRequest {
    int32 id = 1;
}

message User {
    int32 id = 1;
    string name = 2;
}

service UserService {
    rpc GetUser(UserRequest) returns (User);
}
`

func main() {
	fmt.Println("=== Protobuf Definition ===")
	fmt.Print(ProtoFileContent)
	fmt.Println("==========================")
	fmt.Println("実際の開発では 'protoc' コマンドを使ってこのファイルからGoコードを生成します。")
}
