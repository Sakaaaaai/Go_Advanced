package main

import "fmt"

// テーブルを作成するためのSQL (Up Migration)
const UpMigration = `
CREATE TABLE ___ (
	___ SERIAL PRIMARY KEY,
	___ VARCHAR(50) NOT NULL,
	___ TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

// テーブルを削除するためのSQL (Down Migration)
const DownMigration = `
DROP TABLE IF ___ ___;
`

func main() {
	fmt.Println("=== Up Migration SQL ===")
	fmt.Println(UpMigration)
	
	fmt.Println("=== Down Migration SQL ===")
	fmt.Println(DownMigration)
}
