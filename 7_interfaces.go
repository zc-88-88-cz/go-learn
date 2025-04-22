package main

import "fmt"

// 7. 接口
// 接口示例
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    // 1. 接口实现
    rect := Rectangle{Width: 3, Height: 4}
    circle := Circle{Radius: 5}
    
    fmt.Println("接口示例:")
    printShapeInfo(rect)
    printShapeInfo(circle)
    
    // 2. 空接口
    var any interface{}
    any = 10
    fmt.Println("\n空接口值:", any)
    
    any = "Go语言"
    fmt.Println("空接口值:", any)
    
    // 3. 类型断言
    if str, ok := any.(string); ok {
        fmt.Println("字符串长度:", len(str))
    }
    
    // 4. 类型switch
    switch v := any.(type) {
    case int:
        fmt.Println("整数:", v)
    case string:
        fmt.Println("字符串:", v)
    default:
        fmt.Println("未知类型")
    }
}