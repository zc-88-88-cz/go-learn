package main

import (
	"errors"
	"fmt"
	"os"
)

// 错误处理示例
// 8. 错误处理
func main() {
	// 1. 基本错误处理
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	// 2. 自定义错误
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	}

	// 3. 创建新错误
	err = errors.New("这是一个自定义错误")
	fmt.Println("自定义错误:", err)

	// 4. 文件操作错误处理
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Println("文件错误:", err)
		return
	}
	defer file.Close()

	// 5. panic和recover
	fmt.Println("\npanic和recover示例:")
	fmt.Println("调用safeDivide:", safeDivide(10, 0))
}

// 除法函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// 使用panic和recover
func safeDivide(a, b float64) (result float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到panic:", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("除数不能为0")
	}

	return a / b
}