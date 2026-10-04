package main

import (
	"database/sql"
	"testing"
)

func TestSetupDBSignature(t *testing.T) {
	var _ func(*sql.DB) = SetupDB
}
