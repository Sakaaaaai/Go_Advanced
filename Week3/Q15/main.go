package main

import "fmt"

// Protobuf定義ファイルのコンテンツ
const ProtoFileContent = `
// 1. syntaxの指定
syntax = "___";

// UserRequestメッセージの定義（問題要件として既に存在するものとします）
message UserRequest {
    int32 id = 1;
}

// 2. Userメッセージの定義
message ___ {
    // 3. id と name の定義
    ___ ___ = 1;
    ___ ___ = 2;
}

// 4. UserServiceの定義
service ___ {
    // 5. GetUser RPCメソッドの定義
    rpc ___(UserRequest) returns (___);
}
`

func main() {
	fmt.Println("=== Protobuf Definition ===")
	fmt.Println(ProtoFileContent)
	fmt.Println("==========================")
	fmt.Println("実際の開発では 'protoc' コマンドを使ってこのファイルからGoコードを生成します。")
}
