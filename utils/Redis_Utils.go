package utils

import (
	"StudentManger/configs"
	"StudentManger/module"
	"context"
	"fmt"
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
		Addr:     fmt.Sprintf("%s:%s", configs.GetStringConfig("redis.host"), configs.GetStringConfig("redis.port")),
		Password: configs.GetStringConfig("redis.password"), // no password set
		DB:       configs.GetIntConfig("redis.db"),          // use default DB
		PoolSize: 100,                                       // 连接池大小
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
	// 返回指定集合的所有成员
	val := rdb.SMembers(ctx, key)
	return val
}
func SAdd(key string, value interface{}) bool {
	// 向指定集合存入值
	val := rdb.SAdd(ctx, key, value)
	if val.Val() == 1 {
		return true
	}
	return false
}

func SDel(key string, value interface{}) bool {
	// 从集合中删除指定值
	val := rdb.SRem(ctx, key, value)
	if val.Val() == 0 {
		return false
	} else {
		return true
	}
}

func SIsMember(key string, value interface{}) bool {
	// 判断指定value是否在key的集合中
	val := rdb.SIsMember(ctx, key, value)
	return val.Val()
}

func HMGet(key string) map[string]string {
	// 获取某个key下的所有key,value
	val := rdb.HGetAll(ctx, key).Val()
	return val

}
func HMSet(key string, value interface{}) bool {
	// 设置某个key下的key，value
	val := rdb.HMSet(ctx, key, value).Val()
	return val
}
func HMDel(key string, field string) int64 {
	// 删除某个key下的某个字段
	val := rdb.HDel(ctx, key, field).Val()
	return val
}
func HMIsMember(key string, field string) bool {
	val := rdb.HExists(ctx, key, field).Val()
	return val
}

func HMLen(key string) int64 {
	val := rdb.HLen(ctx, key).Val()
	return val
}

func Exists(key string) bool {
	// 判断key是否存在
	val := rdb.Exists(ctx, key)
	if val.Val() == 1 {
		return true
	} else {
		return false
	}
}

func Staff2Map(staff []module.CourseStaff) map[string]string {
	value := make(map[string]string)
	for _, v := range staff {
		value[v.UID] = v.Student
	}
	return value
}

func Del(key string) bool {
	// 删除一个key
	val := rdb.Del(ctx, key)
	if val.Val() == 1 {
		return true
	}
	return false
}
func SDiff(key1 string, key2 string) *redis.StringSliceCmd {
	// 比较两个集合的不同
	val := rdb.SDiff(ctx, key1, key2)
	return val
}
