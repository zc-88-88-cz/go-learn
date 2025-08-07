package main

import "fmt"

// 函数示例
// 2. 函数
func main() {
    // 1. 基本函数调用
    greet("Go语言")
    
    // 2. 多返回值函数
    sum, diff := calculate(10, 5)
    fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
    
    // 3. 命名返回值
    area := rectArea(3, 4)
    fmt.Println("Rectangle area:", area)
    
    // 4. 可变参数函数
    total := sumNumbers(1, 2, 3, 4, 5)
    fmt.Println("Total:", total)
}

// 基本函数
func greet(name string) {
    fmt.Println("Hello,", name)
}

// 多返回值函数
func calculate(a, b int) (int, int) {
    return a + b, a - b
}

// 命名返回值
func rectArea(width, height float64) (area float64) {
    area = width * height // area := width * height  错误的，不能用短变量声明
    return // 自动返回area变量
}

// 可变参数函数
func sumNumbers(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}