package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串操作示例
func main() {
	// 1. 字符串拼接
	fmt.Println("字符串拼接:")
	str1 := "Hello"
	str2 := "世界"
	fmt.Println(str1 + " " + str2)
	fmt.Println(fmt.Sprintf("%s %s", str1, str2))
	fmt.Println(strings.Join([]string{str1, str2}, " "))

	// 2. 字符串分割
	fmt.Println("\n字符串分割:")
	csv := "Go,Python,Java,Rust"
	fmt.Println(strings.Split(csv, ","))
	fmt.Println(strings.SplitAfter(csv, ","))

	// 3. 字符串替换
	fmt.Println("\n字符串替换:")
	replaceStr := "Golang is great, Golang is fast"
	fmt.Println(strings.Replace(replaceStr, "Golang", "Go", 1))
	fmt.Println(strings.ReplaceAll(replaceStr, "Golang", "Go"))

	// 4. 字符串查找
	fmt.Println("\n字符串查找:")
	searchStr := "Searching in this string"
	fmt.Println(strings.Contains(searchStr, "this"))
	fmt.Println(strings.Index(searchStr, "in"))
	fmt.Println(strings.HasPrefix(searchStr, "Search"))
	fmt.Println(strings.HasSuffix(searchStr, "string"))

	// 5. 字符串转换
	fmt.Println("\n字符串转换:")
	numStr := "123"
	num, _ := strconv.Atoi(numStr)
	fmt.Printf("字符串转整数: %d\n", num)
	fmt.Println(strconv.Itoa(456))

	// 6. 字符串修剪
	fmt.Println("\n字符串修剪:")
	trimStr := "   Trim this string   "
	fmt.Printf("|%s|\n", strings.TrimSpace(trimStr))
	fmt.Printf("|%s|\n", strings.Trim(trimStr, " "))
	fmt.Printf("|%s|\n", strings.TrimLeft(trimStr, " "))
	fmt.Printf("|%s|\n", strings.TrimRight(trimStr, " "))

	// 7. 字符串大小写
	fmt.Println("\n字符串大小写:")
	caseStr := "Go Programming"
	fmt.Println(strings.ToUpper(caseStr))
	fmt.Println(strings.ToLower(caseStr))
	fmt.Println(strings.Title(caseStr))

	// 8. 字符串比较
	fmt.Println("\n字符串比较:")
	cmpStr1 := "Go"
	cmpStr2 := "go"
	fmt.Println(cmpStr1 == cmpStr2)
	fmt.Println(strings.EqualFold(cmpStr1, cmpStr2))

	// 9. 字符串缓冲区处理
	fmt.Println("\n字符串缓冲区处理:")
	var buf bytes.Buffer

	// 基本写入操作
	buf.WriteString("Hello ")
	buf.WriteString("世界")
	fmt.Println(buf.String())

	// 重置缓冲区
	buf.Reset()
	buf.WriteString("缓冲区重置后: ")
	buf.WriteByte('A')
	buf.WriteRune('中')
	buf.Write([]byte{" ", 'G', 'o'})
	fmt.Println(buf.String())

	// 高效处理大量数据
	buf.Reset()
	for i := 0; i < 5; i++ {
	    fmt.Fprintf(&buf, "第%d次 ", i+1)
	}
	fmt.Println("循环写入结果:", buf.String())

	// 读取操作
	readBuf := bytes.NewBufferString("读取示例")
	readData, _ := readBuf.ReadByte()
	fmt.Printf("读取第一个字节: %c\n", readData)
	fmt.Println("剩余内容:", readBuf.String())
}