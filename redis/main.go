package main

import (
	"encoding/json"
	"fmt"
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
	_ = LPush("redis-list", "A", "B", 1, 2, 3, "C")
	//str := RPop("redis-list")
	//fmt.Printf("%v", str)
	for {
		ret := BRPop("redis-list")
		fmt.Printf("%v\n", ret)
	}

}
