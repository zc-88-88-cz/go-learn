package main

import "fmt"

// 6. 结构体
// 结构体示例
type Person struct {
    Name string
    Age  int
    Email string
}

func main() {
    // 1. 创建结构体实例
    p1 := Person{"张三", 25, "zhangsan@example.com"}
    
    // 2. 指定字段名创建
    p2 := Person{
        Name: "李四",
        Age:  30,
        Email: "lisi@example.com",
    }
    
    // 3. 匿名结构体
    emp := struct {
        Name string
        Salary float64
    }{
        Name: "王五",
        Salary: 5000.00,
    }
    
    fmt.Println("结构体示例:")
    fmt.Println("p1:", p1)
    fmt.Println("p2:", p2)
    fmt.Println("emp:", emp)
    
    // 4. 访问和修改字段
    p1.Age = 26
    fmt.Println("修改后p1年龄:", p1.Age)
    
    // 5. 结构体指针
    p3 := &Person{"赵六", 40, "zhaoliu@example.com"}
    fmt.Println("p3:", *p3)
    
    // 6. 嵌套结构体
    type Address struct {
        City, Country string
    }
    
    type Employee struct {
        Person
        Address
        Position string
    }
    
    e := Employee{
        Person: Person{"钱七", 35, "qianqi@example.com"},
        Address: Address{"北京", "中国"},
        Position: "工程师",
    }
    
    fmt.Println("\n嵌套结构体:")
    fmt.Println("e:", e)
    fmt.Println("e.Name:", e.Name) // 提升字段
}