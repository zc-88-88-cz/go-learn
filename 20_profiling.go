package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// 1. CPU性能分析
	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("创建CPU性能分析文件失败:", err)
		return
	}
	defer cpuProfile.Close()
	
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()
	
	// 2. 内存性能分析
	memProfile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("创建内存性能分析文件失败:", err)
		return
	}
	defer memProfile.Close()
	
	// 3. 示例性能分析代码
	for i := 0; i < 1000; i++ {
		allocateMemory()
	}
	
	runtime.GC() // 执行GC确保内存统计准确
	pprof.WriteHeapProfile(memProfile)
	
	// 4. 阻塞分析
	blockProfile, err := os.Create("block.prof")
	if err != nil {
		fmt.Println("创建阻塞分析文件失败:", err)
		return
	}
	defer blockProfile.Close()
	
	runtime.SetBlockProfileRate(1) // 记录所有阻塞事件
	defer runtime.SetBlockProfileRate(0)
	
	// 5. 模拟阻塞操作
	doBlockingOperation()
	
	pprof.Lookup("block").WriteTo(blockProfile, 0)
	
	fmt.Println("性能剖析示例完成，请使用go tool pprof分析生成的文件")
}

func allocateMemory() {
	// 分配内存用于内存分析
	data := make([]byte, 1024*1024) // 1MB
	_ = data
	
	// 模拟CPU密集型任务
	for i := 0; i < 1000; i++ {
		_ = i * i
	}
}

func doBlockingOperation() {
	// 模拟阻塞操作
	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(ch)
	}()
	<-ch
}