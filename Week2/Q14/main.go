package main

import (
	"fmt"
	"time"
)

// ========= モック環境 =========
type redisClient struct {
	locks map[string]bool
}
func (r *redisClient) SetNX(key string, value interface{}, expiration time.Duration) bool {
	if r.locks[key] {
		return false // 既にロックされている
	}
	r.locks[key] = true
	return true
}
func (r *redisClient) Del(key string) {
	delete(r.locks, key)
}
// =============================

func ExecuteWithLock(rdb *redisClient, process func()) bool {
	key := "lock:billing"

	// 1. ロックの取得を試みる（SetNX）
	ok := rdb.___(___, "1", ___ * time.Second)
	if !ok {
		// 2. 取得できなければ諦めて false を返す
		fmt.Println("Already running on another instance, skipping...")
		return ___
	}

	// 3. 関数を抜ける際に必ずロックを解放する
	___ func() {
		fmt.Println("Releasing lock...")
		rdb.___(___)
	}()

	// 4. ロックが取れたら処理を実行する
	fmt.Println("Lock acquired! Executing process...")
	___()

	return ___
}

func main() {
	rdb := &redisClient{locks: make(map[string]bool)}

	// ダミーの処理関数
	task := func() {
		fmt.Println("Processing billing data...")
		time.Sleep(100 * time.Millisecond)
	}

	// 1回目は成功する
	success1 := ExecuteWithLock(rdb, task)
	fmt.Println("Execution 1 success:", success1)

	// rdb.locks["lock:billing"] = true // 手動でロック状態を作ってから試すと...
	// success2 := ExecuteWithLock(rdb, task)
	// fmt.Println("Execution 2 success:", success2)
}
