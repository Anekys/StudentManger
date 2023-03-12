package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func init() {
	err := initClient()
	if err != nil {
		panic(err)
	}
}
func SGet(key string) *redis.StringSliceCmd {
	val := rdb.SMembers(ctx, key)
	return val
}
func SAdd(key string, value interface{}) bool {
	val := rdb.SAdd(ctx, key, value)
	if val.Val() == 1 {
		return true
	}
	return false
}
func Del(key string) bool {
	val := rdb.Del(ctx, key)
	if val.Val() == 1 {
		return true
	}
	return false
}

func SDiff(key1 string, key2 string) *redis.StringSliceCmd {
	val := rdb.SDiff(ctx, key1, key2)
	return val
}
