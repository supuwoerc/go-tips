package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	fmt.Println("---redis存入字符串---")
	_ = Set("redis-string", "测试存入redis一段字符串")
	fmt.Println("---redis获取字符串---")
	result, _ := Get("redis-string")
	fmt.Println(result)
	fmt.Println("---redis存入结构体---")
	var user = User{
		Name: "测试结构体名字",
		Age:  19,
	}
	jsonString, _ := json.Marshal(user)
	_ = Set("user-struct", jsonString)
	fmt.Println("---redis取出结构体---")
	var userResult User
	getResult, _ := Get("user-struct")
	_ = json.Unmarshal([]byte(getResult), &userResult)
	fmt.Println(userResult)
	fmt.Println("---redis操作list---")
	_ = LPush("redis-list1", "A", "B", 1, 2, 3, "C")
	_ = LPush("redis-list2", "A", "B", 1, 2, 3, "C")
	str := RPop("redis-list1")
	fmt.Printf("%v", str)
	length, _ := LLen("redis-list1")
	fmt.Println(length)
	lRange, _ := LRange("redis-list1", 0, 3)
	fmt.Println(lRange)
	fmt.Println("------")
	for {
		ret, err := BRPop("redis-list2")
		if err != nil {
			break
		}
		fmt.Printf("%v\n", ret[1])
	}
	hashKey := "hash-key"
	setResult, err := HSet(hashKey, "name", "测试名称")
	fmt.Println(setResult, err)
	if err != nil {
		fmt.Println(err)
	} else {
		get, _ := HGet(hashKey, "name")
		fmt.Println(get)
	}
	exists, _ := HExists(hashKey, "name")
	fmt.Println(exists)
	all, _ := HGetAll(hashKey)
	fmt.Println(all)
	setKey := "set-key"
	_, err = SAdd(setKey, 1, 3, 4, 5, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		length, _ := SCard(setKey)
		fmt.Println(length)
		fmt.Println(SMembers(setKey))
	}
	zSetKey := "z-set"

	_, err = ZAdd(zSetKey, redis.Z{
		Score:  0,
		Member: "零分",
	}, redis.Z{
		Score:  3,
		Member: "三分",
	}, redis.Z{
		Score:  2,
		Member: "两分",
	}, redis.Z{
		6,
		"六分",
	}, redis.Z{
		Score:  8,
		Member: "八分",
	})
	if err == nil {
		scores, _ := ZRevRangeWithScores(zSetKey, 0, 2)
		fmt.Println(scores)
		_, _ = ZIncrBy(zSetKey, 10, "零分")
		fmt.Println(ZRevRangeWithScores(zSetKey, 0, 2))
	}
}
