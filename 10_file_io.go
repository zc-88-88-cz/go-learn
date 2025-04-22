package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 文件操作示例
func main() {
	// 1. 创建文件
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. 写入文件
	content := []byte("Hello, 这是文件操作示例!")
	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}

	// 3. 读取文件
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("文件内容:", string(data))

	// 4. 文件信息
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("文件名: %s, 大小: %d bytes\n", fileInfo.Name(), fileInfo.Size())

	// 5. 删除文件
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("文件已删除")
}