//go:build ignore

package main

import "fmt"

const UpMigration = `
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const DownMigration = `
DROP TABLE IF EXISTS users;
`

func main() {
	fmt.Println("=== Up Migration SQL ===")
	fmt.Print(UpMigration)
	
	fmt.Println("=== Down Migration SQL ===")
	fmt.Print(DownMigration)
}
