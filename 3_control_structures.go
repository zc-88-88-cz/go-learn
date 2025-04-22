package main

import "fmt"

// 3. 控制结构
// 控制结构示例
func main() {
    // 1. if-else语句
    age := 18
    if age >= 18 {
        fmt.Println("成年人")
    } else {
        fmt.Println("未成年人")
    }
    
    // 2. for循环
    fmt.Println("\nfor循环示例:")
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", i)
    }
    
    // 3. while循环(Go中没有while关键字，用for代替)
    fmt.Println("\n\nwhile循环示例:")
    j := 0
    for j < 3 {
        fmt.Printf("%d ", j)
        j++
    }
    
    // 4. 无限循环
    /*
    for {
        fmt.Println("无限循环")
    }
    */
    
    // 5. switch语句
    fmt.Println("\n\nswitch语句示例:")
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("星期一")
    case "Tuesday":
        fmt.Println("星期二")
    default:
        fmt.Println("其他日子")
    }
    
    // 6. 带条件的switch
    score := 85
    switch {
    case score >= 90:
        fmt.Println("优秀")
    case score >= 80:
        fmt.Println("良好")
    case score >= 60:
        fmt.Println("及格")
    default:
        fmt.Println("不及格")
    }
}