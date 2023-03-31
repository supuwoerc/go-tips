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

// HSet 设置hash键值对
func HSet(key, field string, value any) (bool, error) {
	result, err := RedisCache.HSet(key, field, value).Result()
	return result, err
}

// HGet 获取hash到键对应到值
func HGet(key, field string) (any, error) {
	result, err := RedisCache.HGet(key, field).Result()
	return result, err
}

// HExists 判断hash存不存在对应到key
func HExists(key, field string) (bool, error) {
	result, err := RedisCache.HExists(key, field).Result()
	return result, err
}

// HGetAll 获取全部entries
func HGetAll(key string) (map[string]string, error) {
	result, err := RedisCache.HGetAll(key).Result()
	return result, err
}

// HDel 删除hash对应到键值对
func HDel(key, field string) (int64, error) {
	result, err := RedisCache.HDel(key, field).Result()
	return result, err
}

// SAdd Set添加元素
func SAdd(key string, members ...any) (int64, error) {
	result, err := RedisCache.SAdd(key, members...).Result()
	return result, err
}

// SCard Set查询成员数量
func SCard(key string) (int64, error) {
	result, err := RedisCache.SCard(key).Result()
	return result, err
}

// SMembers Set查询全部成员
func SMembers(key string) ([]string, error) {
	result, err := RedisCache.SMembers(key).Result()
	return result, err
}

// ZAdd zSet添加元素
func ZAdd(key string, members ...redis.Z) (int64, error) {
	result, err := RedisCache.ZAdd(key, members...).Result()
	return result, err
}

// ZRevRangeWithScores zSet按照权重返回数据
func ZRevRangeWithScores(key string, start, end int64) ([]redis.Z, error) {
	zs, err2 := RedisCache.ZRevRangeWithScores(key, start, end).Result()
	return zs, err2
}

// ZIncrBy 添加权重
func ZIncrBy(key string, increment float64, member string) (float64, error) {
	result, err := RedisCache.ZIncrBy(key, increment, member).Result()
	return result, err
}

// Expire 设置相对过期时间
func Expire(key string, duration time.Duration) (bool, error) {
	result, err := RedisCache.Expire(key, duration).Result()
	return result, err
}

// ExpireAt 设置绝对过期时间
func ExpireAt(key string, tm time.Time) (bool, error) {
	result, err := RedisCache.ExpireAt(key, tm).Result()
	return result, err
}
