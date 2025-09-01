// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
package main

import (
	"fmt"
	"time"
)

// sender 协程：生成数字并发送到通道
func sender(ch chan<- int) {
	fmt.Println("发送协程开始工作...")
	for i := 1; i <= 10; i++ {
		fmt.Printf("发送数字: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	fmt.Println("发送协程工作完成")
}

// receiver 协程：从通道接收数字并打印
func receiver(ch <-chan int) {
	fmt.Println("接收协程开始工作...")
	for num := range ch {
		fmt.Printf("接收到数字: %d\n", num)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("接收协程工作完成")
}

func main() {
	// 创建一个无缓冲通道
	ch := make(chan int)

	// 启动发送协程
	go sender(ch)

	// 启动接收协程
	go receiver(ch)
	time.Sleep(1 * time.Second)
}
