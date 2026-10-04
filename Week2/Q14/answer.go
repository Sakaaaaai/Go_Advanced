//go:build ignore

package main

import (
	"fmt"
	"time"
)

type redisClient struct {
	locks map[string]bool
}
func (r *redisClient) SetNX(key string, value interface{}, expiration time.Duration) bool {
	if r.locks[key] {
		return false
	}
	r.locks[key] = true
	return true
}
func (r *redisClient) Del(key string) {
	delete(r.locks, key)
}

func ExecuteWithLock(rdb *redisClient, process func()) bool {
	key := "lock:billing"

	ok := rdb.SetNX(key, "1", 5 * time.Second)
	if !ok {
		fmt.Println("Already running on another instance, skipping...")
		return false
	}

	defer func() {
		fmt.Println("Releasing lock...")
		rdb.Del(key)
	}()

	fmt.Println("Lock acquired! Executing process...")
	process()

	return true
}

func main() {
	rdb := &redisClient{locks: make(map[string]bool)}

	task := func() {
		fmt.Println("Processing billing data...")
		time.Sleep(100 * time.Millisecond)
	}

	success1 := ExecuteWithLock(rdb, task)
	fmt.Println("Execution 1 success:", success1)
}
