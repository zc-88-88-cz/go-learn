package main

import (
	"fmt"
	"reflect"
)

// 反射机制示例
type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. 基本类型反射
	var x float64 = 3.14
	fmt.Println("类型:", reflect.TypeOf(x))
	fmt.Println("值:", reflect.ValueOf(x))

	// 2. 结构体反射
	p := Person{"张三", 25}
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("字段 %d: %s = %v\n", i, t.Field(i).Name, v.Field(i).Interface())
	}

	// 3. 修改值
	px := &Person{"李四", 30}
	v = reflect.ValueOf(px).Elem()
	f := v.FieldByName("Age")
	if f.CanSet() {
		f.SetInt(35)
		fmt.Println("修改后的年龄:", px.Age)
	}
}