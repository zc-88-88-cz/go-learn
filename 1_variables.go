package main

import "fmt"
import "os"
import "bytes"

// 变量声明示例
// 1. 变量
func main() {
    // 1. 使用var声明变量
    var name string = "Go语言"
    // var version2 string = "1.18"
    // 2. 类型推断
    var version = 1.18
	// var age3 = 10
    
    // 3. 短变量声明(只能在函数内使用)
    isAwesome := true
	// haha := "hahahha nihao"
    
    // 4. 多变量声明
    var (
        age int = 10
        height float64 = 1.75
    )
    
    // 5. 常量声明
    const pi = 3.1415926
    //const pai := 3.1415926 // 错误的，常量不能用短 声明 :=
	const name333 string = "Go语言"

	const (
		language = "Go"
		area = "China"
		Type = 1
	)
    fmt.Println("变量声明示例:")
    fmt.Println("name:", name)
    fmt.Println("version:", version)
    fmt.Println("isAwesome:", isAwesome)
    fmt.Println("age:", age, "height:", height)
    fmt.Println("pi:", pi)

	fmt.Print("直接打印", "不换行", 123, 77777)
	fmt.Print(1,4,8,9)
	fmt.Println(1,4,8,9)
	fmt.Println(1,4,8,9)
	fmt.Printf("格式化打印 %s %d %f", "hahah", 123, 1.23)
	fmt.Printf("%v", "------")  // 默认格式
	fmt.Printf("%+v", 33333) // 结构体字段名+值 
	fmt.Println("--------")
	fmt.Printf("%#v", true) // Go语法表示
	fmt.Println("--------")
	fmt.Printf("%d", 100) // 十进制
	fmt.Println("--------")
	fmt.Printf("%f", 12.345) // 浮点数
	fmt.Println("--------")
	fmt.Printf("%s", "332232") // 字符串表示
	fmt.Println("--------")
	fmt.Printf("%p", &version) // 指针地址
	fmt.Println("--------")
	fmt.Printf("|%4d|%4d|\n", 2, 44) // // 数字右对齐 宽度为 4
	fmt.Printf("|%-4d|%-4d|\n", 2, 44) // 数字左对齐 宽度为 4
	fmt.Printf("|%10.2f|%10.2f|\n", 1333.2, 3333.457) // 浮点数 宽度为 4 精度为 2


	type Point struct {
		X, Y int
	}
	p := Point{1, 2}
	fmt.Printf("%v\n", p) // {1 2}
	fmt.Printf("%+v\n", p) // {X:1 Y:2}
	fmt.Printf("%#v\n", p) // main.Point{X:1, Y:2}
	fmt.Printf("%T\n", p) // main.Point
	fmt.Printf("%t\n", true) // true

	_, err := os.Open("不存在的文件")
	if err != nil {
		fmt.Printf("错误: %v\n", err) // 错误: open 不存在的文件: no such file or directory
	}

	var name6 string
	var age6 int

	fmt.Scan(&name6, &age6)
	fmt.Printf("姓名: %s, 年龄: %d\n", name6, age6)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "姓名: %s, 年龄: %d\n", name6, age6)
	fmt.Println(buf.String())

	// 表格对齐输出
	fmt.Printf("|%-10s|%10s|%10s|\n", "产品", "价格", "库存")
	fmt.Printf("|%-10s|%10.2f|%10d|\n", "手机", 2999.99, 50)
	fmt.Printf("|%-10s|%10.2f|%10d|\n", "笔记本", 5999.00, 20)

	// 错误信息格式化
	err11 := fmt.Errorf("参数错误: 期望 %d 得到 %d", 10, 5)
	fmt.Printf("错误信息: %v\n", err11)
}