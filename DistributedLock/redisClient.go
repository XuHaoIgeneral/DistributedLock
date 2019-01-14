package DistributedLock

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

func Init() {
	host := ""
	port := ""
	db := 0
	passWord := ""
	var addr = fmt.Sprintf("%s:%s", host, port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passWord, // no password set
		DB:       int(db),  // use default DB
	})

	//测试redis数据
	err := RedisClient.Set("test", "test", 30*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
