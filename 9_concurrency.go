package main

import (
	"fmt"
	"time"
)

// 9. 并发
// 并发编程示例
func main() {
	// 1. goroutine基础
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	// 2. 使用channel通信
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println("收到消息:", msg)

	// 3. channel缓冲
	buffered := make(chan string, 2)
	buffered <- "缓冲"
	buffered <- "channel"
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	// 4. channel同步
	done := make(chan bool)
	go worker(done)
	<-done

	// 5. select多路复用
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("收到:", msg1)
		case msg2 := <-c2:
			fmt.Println("收到:", msg2)
		}
	}
}

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}