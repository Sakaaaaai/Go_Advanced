package main

import "testing"

func TestGORMFunctionsSignature(t *testing.T) {
	var _ func(*gormDB, string, int) error = CreateUser
	var _ func(*gormDB, string) (User, error) = GetUserByName
}
