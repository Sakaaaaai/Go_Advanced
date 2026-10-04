//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"time"
)

func SetupDB(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
}

func main() {
	fmt.Println("This is a package for DB configuration logic.")
}
