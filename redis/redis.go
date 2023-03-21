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

// LPush 存储数据到list
func LPush(key string, value ...interface{}) error {
	_, err := RedisCache.LPush(key, value...).Result()
	return err
}

// RPop list取出数据
func RPop(key string) interface{} {
	result, _ := RedisCache.RPop(key).Result()
	return result
}

// BRPop list取出数据,堵塞
func BRPop(key string) ([]string, error) {
	result, err := RedisCache.BRPop(3*time.Second, key).Result()
	return result, err
}

// LLen list长度
func LLen(key string) (int64, error) {
	length, err := RedisCache.LLen(key).Result()
	return length, err
}

// LRange list长度
func LRange(key string, s, e int64) ([]string, error) {
	result, err := RedisCache.LRange(key, s, e).Result()
	return result, err
}
