package main

import (
	"fmt"
	"plugin"
)

func main() {
	// 1. 加载插件
	p, err := plugin.Open("plugin_example.so")
	if err != nil {
		fmt.Println("Error loading plugin:", err)
		return
	}
	
	// 2. 查找导出的符号
	// 查找函数
	funcSym, err := p.Lookup("Greet")
	if err != nil {
		fmt.Println("Error looking up function:", err)
		return
	}
	
	// 类型断言为函数
	greet, ok := funcSym.(func(string) string)
	if !ok {
		fmt.Println("Unexpected type from module symbol")
		return
	}
	
	// 3. 调用插件函数
	fmt.Println("Plugin says:", greet("Go Developer"))
	
	// 4. 查找变量
	varSym, err := p.Lookup("Message")
	if err != nil {
		fmt.Println("Error looking up variable:", err)
		return
	}
	
	// 类型断言为字符串指针
	message, ok := varSym.(*string)
	if !ok {
		fmt.Println("Unexpected type from module symbol")
		return
	}
	
	fmt.Println("Plugin variable:", *message)
	
	fmt.Println("插件系统示例完成")
}

/*
插件示例代码 (需要单独编译为plugin_example.so):

package main

import "C"

var Message = "Hello from plugin!"

//export Greet
func Greet(name string) string {
	return "Hello, " + name + "!"
}

func main() {}
*/