package main

import "fmt"

// 5. map
// map示例
func main() {
    // 1. 创建map
    m1 := make(map[string]int)
    m1["apple"] = 5
    m1["banana"] = 7
    
    // 2. 字面量创建map
    m2 := map[string]int{
        "orange": 3,
        "pear":  4,
    }
    
    fmt.Println("map示例:")
    fmt.Println("m1:", m1)
    fmt.Println("m2:", m2)
    
    // 3. 访问和修改map
    fmt.Println("\n访问和修改map:")
    fmt.Println("apple数量:", m1["apple"])
    m1["apple"] = 10
    fmt.Println("修改后apple数量:", m1["apple"])
    
    // 4. 检查key是否存在
    if value, ok := m1["grape"]; ok {
        fmt.Println("grape数量:", value)
    } else {
        fmt.Println("grape不存在")
    }
    
    // 5. 删除key
    delete(m1, "banana")
    fmt.Println("删除banana后m1:", m1)
    
    // 6. 遍历map
    fmt.Println("\n遍历map:")
    for key, value := range m2 {
        fmt.Printf("水果:%s 数量:%d\n", key, value)
    }
}