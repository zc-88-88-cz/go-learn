package main

/*
#include <stdio.h>
#include <stdlib.h>

// C函数声明
void helloFromC() {
    printf("Hello from C!\n");
}

int add(int a, int b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
    // 1. 调用简单C函数
    fmt.Println("Calling C function from Go:")
    C.helloFromC()
    
    // 2. 调用带参数的C函数
    sum := C.add(3, 5)
    fmt.Printf("C add result: %d\n", sum)
    
    // 3. 使用C标准库
    cstr := C.CString("Go string to C")
    defer C.free(unsafe.Pointer(cstr)) // 释放C分配的内存
    
    // 4. 复杂类型转换示例
    goSlice := []byte{'G', 'o', ' ', 'a', 'n', 'd', ' ', 'C'}
    cArr := (*C.char)(unsafe.Pointer(&goSlice[0]))
    length := C.int(len(goSlice))
    
    // 5. 调用C函数处理字节数组
    processBytes(cArr, length)
    
    fmt.Println("CGO示例完成")
}

// 导出Go函数给C调用
//export processBytes
export func processBytes(data *C.char, length C.int) {
    // C代码可以调用此函数
    fmt.Printf("Processing %d bytes from C\n", length)
}