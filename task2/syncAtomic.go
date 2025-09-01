// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64

	fmt.Println("开始原子操作计数...")
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(goroutineID int) {
			defer wg.Done()
			for j := 1; j <= 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}(i)
	}
	wg.Wait()

	elapsedTime := time.Since(startTime)

	// 输出结果
	fmt.Printf("\n执行时间: %v\n", elapsedTime)
	fmt.Printf("期望值: %d\n", 10*1000)
	fmt.Printf("实际值: %d\n", atomic.LoadInt64(&counter))

	if atomic.LoadInt64(&counter) == 10*1000 {
		fmt.Println("✅ 计数正确，原子操作保证了数据安全")
	} else {
		fmt.Println("❌ 计数错误")
	}
}
