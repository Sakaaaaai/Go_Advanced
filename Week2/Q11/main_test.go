package main

import "testing"

func TestEagerLoadingSignature(t *testing.T) {
	var _ func(*gormDB) []User = GetUsersEager
}
