//go:build ignore

package main

import (
	"fmt"
	"time"
)

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
	return "Profile of " + key, nil
}

func GetUserProfile(rdb *redisClient, db *dbClient, userID string) (string, error) {
	key := "user:profile:" + userID

	val, err := rdb.Get(key).Result()
	if err == nil {
		fmt.Println("Cache Hit!")
		return val, nil
	}

	fmt.Println("Cache Miss! Fetching from DB...")
	profile, err := db.Get(userID)
	if err != nil {
		return "", err
	}

	rdb.Set(key, profile, 10 * time.Minute)

	return profile, nil
}

func main() {
	rdb := &redisClient{data: make(map[string]string)}
	db := &dbClient{}

	fmt.Println(GetUserProfile(rdb, db, "123"))
	fmt.Println(GetUserProfile(rdb, db, "123"))
}
