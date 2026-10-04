package main

import (
	"fmt"
	"time"
)

// ========= モック環境 =========
type stringCmd struct {
	val string
	err error
}
func (cmd *stringCmd) Result() (string, error) { return cmd.val, cmd.err }

type redisClient struct {
	data map[string]string
}
func (r *redisClient) Get(key string) *stringCmd {
	val, ok := r.data[key]
	if !ok { return &stringCmd{err: fmt.Errorf("redis: nil")} }
	return &stringCmd{val: val}
}
func (r *redisClient) Set(key string, value interface{}, expiration time.Duration) *stringCmd {
	r.data[key] = fmt.Sprint(value)
	return &stringCmd{val: "OK"}
}

type dbClient struct{}
func (db *dbClient) Get(key string) (string, error) {
	// DBからの取得は重い処理（今回はダミーデータを返す）
	return "Profile of " + key, nil
}
// =============================

func GetUserProfile(rdb *redisClient, db *dbClient, userID string) (string, error) {
	key := "user:profile:" + userID

	// 1. まずRedis（キャッシュ）から取得を試みる
	val, err := rdb.___(___).Result()
	if err == nil {
		// キャッシュヒット
		fmt.Println("Cache Hit!")
		return val, nil
	}

	// 2. キャッシュに無ければDBから取得する
	fmt.Println("Cache Miss! Fetching from DB...")
	profile, err := db.___(___)
	if err != nil {
		return "", err
	}

	// 3. 次回のためにDBから取ったデータをRedisにSetする（有効期限は10分）
	rdb.___(___, profile, ___ * time.Minute)

	return profile, nil
}

func main() {
	rdb := &redisClient{data: make(map[string]string)}
	db := &dbClient{}

	fmt.Println(GetUserProfile(rdb, db, "123")) // 1回目はDBから
	fmt.Println(GetUserProfile(rdb, db, "123")) // 2回目はキャッシュから
}
