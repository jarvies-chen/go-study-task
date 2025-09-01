// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter 计数器结构体
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment 递增计数器
func Increment(c *Counter) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue 获取计数器值
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// SetValue 设置计数器值
func (c *Counter) SetValue(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = value
}

func main() {
	// 创建共享计数器
	counter := &Counter{}

	fmt.Println("开始并发计数...")
	startTime := time.Now()

	var wg sync.WaitGroup

	// 启动10个协程
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			// 每个协程执行1000次递增操作
			for j := 1; j <= 1000; j++ {
				Increment(counter)
			}
			fmt.Printf("协程 %d 完成\n", goroutineID)
		}(i)
	}

	wg.Wait()
	elapsedTime := time.Since(startTime)
	// 输出结果
	fmt.Printf("\n执行时间: %v\n", elapsedTime)
	fmt.Printf("期望值: %d\n", 10*1000)
	fmt.Printf("实际值: %d\n", counter.GetValue())

	if counter.GetValue() == 10*1000 {
		fmt.Println("✅ 计数正确，没有数据竞争")
	} else {
		fmt.Println("❌ 计数错误，存在数据竞争")
	}
}
