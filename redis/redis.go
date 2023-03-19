package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var RedisCache *redis.Client

func init() {
	fmt.Println("redis初始化开始")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	RedisCache = client
	result, err := RedisCache.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis初始化成功", result)
}

// Get 根据key获取string数据
func Get(key string) (string, error) {
	result, err := RedisCache.Get(key).Result()
	return result, err
}

// Set 存储数据到redis,过期时间为24小时
func Set(key string, value interface{}) error {
	_, err := RedisCache.Set(key, value, time.Hour*24).Result()
	return err
}
