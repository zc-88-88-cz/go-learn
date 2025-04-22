package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// 标准库使用示例
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 1. 随机数生成
	rand.Seed(time.Now().UnixNano())
	fmt.Println("随机数:", rand.Intn(100))

	// 2. JSON编码与解码
	user := User{"王五", 28}
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON编码:", string(jsonData))

	var decodedUser User
	err = json.Unmarshal(jsonData, &decodedUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON解码: %+v\n", decodedUser)

	// 3. 时间处理
	now := time.Now()
	fmt.Println("当前时间:", now.Format("2006-01-02 15:04:05"))
	future := now.AddDate(0, 1, 0)
	fmt.Println("一个月后:", future.Format("2006-01-02 15:04:05"))
}