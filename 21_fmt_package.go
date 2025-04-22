package main

import "fmt"

// fmt包功能示例
func main() {
    // 1. 基本格式化输出
    fmt.Println("基本输出:")
    fmt.Println("Hello, 世界")
    fmt.Print("Print不换行 ")
    fmt.Print("继续在同一行")
    fmt.Println() // 换行

    // 2. 格式化输出
    fmt.Println("\n格式化输出:")
    name := "Go语言"
    version := 1.21
    fmt.Printf("名称: %s, 版本: %.2f\n", name, version)

    // 3. 字符串扫描
    fmt.Println("\n字符串扫描:")
    input := "42 3.14 Golang"
    var i int
    var f float64
    var s string
    fmt.Sscan(input, &i, &f, &s)
    fmt.Printf("解析结果: %d %f %s\n", i, f, s)

    // 4. 错误处理
    fmt.Println("\n错误处理:")
    err := fmt.Errorf("这是一个自定义错误: %s", "文件未找到")
    fmt.Println("错误信息:", err)

    // 5. 高级格式化
    fmt.Println("\n高级格式化:")
    fmt.Printf("二进制: %b, 八进制: %o, 十六进制: %x\n", 255, 255, 255)
    fmt.Printf("科学计数法: %e\n", 123456789.0)
    fmt.Printf("字符: %c\n", 65) // ASCII 65 = 'A'

    // 6. 宽度和精度控制
    fmt.Println("\n宽度和精度:")
    fmt.Printf("|%10s|%10s|\n", "产品", "价格")
    fmt.Printf("|%10s|%10.2f|\n", "Go书", 99.99)
    fmt.Printf("|%-10s|%-10.2f|\n", "Go书", 99.99) // 左对齐

    // 7. 字符串格式化
    fmt.Println("\n字符串格式化:")
    fmt.Println(fmt.Sprint("直接拼接", 123, true))
    fmt.Println(fmt.Sprintf("格式化字符串: %.2f", 3.14159))
}