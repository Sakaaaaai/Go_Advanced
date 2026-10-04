package main

import (
	"fmt"
)

// 1. 使い捨てDB（Testcontainers）のセットアップを行う疑似コード
// DBのURLと、破棄用の関数を返します
func setupTestDB() (string, func(), error) {
	fmt.Println("[Testcontainers] Dockerコンテナ（PostgreSQL）を起動中...")
	dbURL := "postgres://testuser:testpass@localhost:5432/testdb"
	
	// 2. コンテナを終了・破棄するためのクリーンアップ関数を定義
	cleanup := func() {
		fmt.Println("[Testcontainers] テスト完了。コンテナを破棄します。")
	}

	return dbURL, cleanup, nil
}

// 3. 結合テスト本体（今回は main 関数でシミュレートしています）
func TestUserRepositoryIntegration() {
	// テスト用のクリーンなDBを立ち上げる
	dbURL, cleanup, err := ___()
	if err != nil {
		fmt.Println("Failed to start container:", err)
		return
	}
	
	// 4. テスト関数を抜けるときに必ずコンテナを破棄する
	___ ___()

	fmt.Printf("接続先DB: %s\n", dbURL)
	
	// 5. 以下、この使い捨てのDBに対して、
	// テーブル作成 -> データのInsert -> データのSelect（検証）を行います
	fmt.Println("Running integration tests...")
	fmt.Println(" -> CREATE TABLE users...")
	fmt.Println(" -> INSERT INTO users ('Alice')...")
	fmt.Println(" -> Check if Alice exists... OK!")
}

func main() {
	fmt.Println("=== 結合テスト開始 ===")
	TestUserRepositoryIntegration()
	fmt.Println("=== 結合テスト終了 ===")
}
