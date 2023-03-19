package main

import (
	"fmt"
)

func main() {
	fmt.Println("---redis存入字符串---")
	_ = Set("redis-string", "redis存入字符串---test")
	fmt.Println("---redis获取字符串---")
	result, _ := Get("redis-string")
	fmt.Println(result)
}
