package main

import "fmt"

// 变量声明示例
// 1. 变量
func main() {
    // 1. 使用var声明变量
    var name string = "Go语言"
    
    // 2. 类型推断
    var version = 1.18
    
    // 3. 短变量声明(只能在函数内使用)
    isAwesome := true
    
    // 4. 多变量声明
    var (
        age int = 10
        height float64 = 1.75
    )
    
    // 5. 常量声明
    const pi = 3.1415926
    
    fmt.Println("变量声明示例:")
    fmt.Println("name:", name)
    fmt.Println("version:", version)
    fmt.Println("isAwesome:", isAwesome)
    fmt.Println("age:", age, "height:", height)
    fmt.Println("pi:", pi)
}