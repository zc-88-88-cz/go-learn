package main

import "fmt"

// 4. 数组和切片
// 数组和切片示例
func main() {
    // 1. 数组声明
    var arr1 [3]int = [3]int{1, 2, 3}
	var tarr1 [4]string = [4]string{"1", "2", "3", "4"}
    arr2 := [...]int{4, 5, 6} // 自动计算长度
	arr3 := [3]int{1: 2, 2: 3} // 索引初始化
	arr4 := [...]int{1,2,3,4,5,6,7,8,9,10}
    
    fmt.Println("数组示例:")
    fmt.Println("arr1:", arr1)
    fmt.Println("arr2:", arr2)
    
    // 2. 切片声明
	var slice0 []int  // 初始化为nil，长度和容量都是0
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := make([]int, 3, 5) // 长度3，容量5

    
    fmt.Println("\n切片示例:")
    fmt.Println("slice1:", slice1, "长度:", len(slice1), "容量:", cap(slice1))
    fmt.Println("slice2:", slice2, "长度:", len(slice2), "容量:", cap(slice2))
    
    // 3. 切片操作
    fmt.Println("\n切片操作:")
    fmt.Println("slice1[1:3]:", slice1[1:3]) // 截取
    
    slice1 = append(slice1, 6) // 追加
    fmt.Println("追加后的slice1:", slice1)
    
    // 4. 遍历切片
    fmt.Println("\n遍历切片:")
    for i, v := range slice1 {
        fmt.Printf("索引:%d 值:%d\n", i, v)
    }
}