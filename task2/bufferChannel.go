// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
package main

import (
	"fmt"
	"sync"
	"time"
)

// producer 生产者协程：向缓冲通道发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("生产者开始工作...")
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("生产者发送: %d\n", i)

		// 模拟生产耗时
		if i%10 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}
	close(ch)
	fmt.Println("生产者工作完成")
}

// consumer 消费者协程：从缓冲通道接收整数
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("消费者开始工作...")
	count := 0

	for num := range ch {
		count++
		fmt.Printf("消费者接收: %d\n", num)

		// 模拟消费耗时
		if count%15 == 0 {
			time.Sleep(150 * time.Millisecond)
		}
	}
	fmt.Printf("消费者总共接收了 %d 个数字\n", count)
	fmt.Println("消费者工作完成")
}

func main() {
	// 创建缓冲通道，缓冲区大小为10
	bufferChan := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)

	// 启动生产者协程
	go producer(bufferChan, &wg)
	// 启动消费者协程
	go consumer(bufferChan, &wg)

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有协程执行完成")
}
