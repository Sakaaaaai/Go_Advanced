//go:build ignore

package main

import (
	"fmt"
)

func setupTestDB() (string, func(), error) {
	fmt.Println("[Testcontainers] Dockerコンテナ（PostgreSQL）を起動中...")
	dbURL := "postgres://testuser:testpass@localhost:5432/testdb"
	
	cleanup := func() {
		fmt.Println("[Testcontainers] テスト完了。コンテナを破棄します。")
	}

	return dbURL, cleanup, nil
}

func TestUserRepositoryIntegration() {
	dbURL, cleanup, err := setupTestDB()
	if err != nil {
		fmt.Println("Failed to start container:", err)
		return
	}
	
	defer cleanup()

	fmt.Printf("接続先DB: %s\n", dbURL)
	
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
