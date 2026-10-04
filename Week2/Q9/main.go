package main

import (
	"database/sql"
	"fmt"
	"time"
)

// SetupDB はデータベースのコネクションプールを最適化します
func SetupDB(db *sql.___) {
	// 1. 最大接続数を 25 に設定する
	db.___(___)

	// 2. 最大アイドル接続数を 25 に設定する
	db.___(___)

	// 3. 接続の最大生存期間を 5分 に設定する
	db.___(___ * time.Minute)
}

func main() {
	fmt.Println("This is a package for DB configuration logic.")
}
