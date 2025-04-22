package main

import (
	"fmt"
	"sync"
	"time"
)

// 高级并发模式示例
func main() {
	// 1. Worker Pool模式
	workerPoolDemo()
	
	// 2. 扇出/扇入模式
	fanOutFanInDemo()
	
	// 3. 并发控制模式
	concurrencyControlDemo()
}

// Worker Pool模式示例
func workerPoolDemo() {
	const numJobs = 10
	const numWorkers = 3
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	// 创建工作池
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	
	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 等待所有worker完成
	wg.Wait()
	close(results)
	
	// 收集结果
	for result := range results {
		fmt.Println("Worker result:", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // 模拟耗时任务
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

// 扇出/扇入模式示例
func fanOutFanInDemo() {
	in := gen(2, 3)
	
	// 扇出 - 将工作分配给多个goroutine
	c1 := sq(in)
	c2 := sq(in)
	
	// 扇入 - 合并多个通道的结果
	for n := range merge(c1, c2) {
		fmt.Println("Fan-out/Fan-in result:", n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	
	// 为每个输入通道启动一个输出goroutine
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	
	// 启动一个goroutine在所有output goroutine完成后关闭out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 并发控制模式示例
func concurrencyControlDemo() {
	// 1. 使用带缓冲的通道限制并发数
	const maxConcurrent = 3
	sem := make(chan struct{}, maxConcurrent)
	
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{} // 获取信号量
			defer func() { <-sem }() // 释放信号量
			
			fmt.Printf("Task %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Task %d finished\n", id)
		}(i)
	}
	wg.Wait()
	
	// 2. 使用context实现取消和超时控制
	// (实际代码中需要添加context包)
	fmt.Println("并发控制示例完成")
}