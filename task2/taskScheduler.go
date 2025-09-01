// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 任务结构体
type Task struct {
	ID       string
	Function func() error
	Result   error
	Duration time.Duration
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks []*Task
}

// NewTaskScheduler 创建任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]*Task, 0),
	}
}

// AddTask 添加任务
func (ts *TaskScheduler) AddTask(id string, fn func() error) {
	task := &Task{
		ID:       id,
		Function: fn,
	}
	ts.tasks = append(ts.tasks, task)
}

// Execute 并发执行所有任务
func (ts *TaskScheduler) Execute() []*Task {
	var wg sync.WaitGroup
	results := make([]*Task, len(ts.tasks))

	//并发所有任务
	for i, task := range ts.tasks {
		wg.Add(1)
		go func(index int, t *Task) {
			defer wg.Done()

			start := time.Now()
			result := t.Function()
			duration := time.Since(start)

			// 保存结果
			results[index] = &Task{
				ID:       t.ID,
				Result:   result,
				Duration: duration,
			}
		}(i, task)
	}
	//等待所有任务完成
	wg.Wait()
	return results
}

// 示例任务函数
func task1() error {
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 Completed")
	return nil
}

func task2() error {
	time.Sleep(1 * time.Second)
	fmt.Println("Task 2 completed")
	return nil
}

func task3() error {
	time.Sleep(3 * time.Second)
	fmt.Println("Task 3 completed")
	return fmt.Errorf("task 3 failed")
}

func main() {
	// 创建调度器
	scheduler := NewTaskScheduler()

	// 添加任务
	scheduler.AddTask("task1", task1)
	scheduler.AddTask("task2", task2)
	scheduler.AddTask("task3", task3)

	fmt.Println("开始执行任务...")
	startTime := time.Now()

	// 执行任务
	results := scheduler.Execute()

	totalTime := time.Since(startTime)
	fmt.Printf("\n所有任务执行完成，总耗时: %v\n", totalTime)

	// 打印结果
	fmt.Println("\n任务执行详情:")
	for _, result := range results {
		if result.Result != nil {
			fmt.Printf("任务 %s: 失败 (%v), 执行时间: %v\n",
				result.ID, result.Result, result.Duration)
		} else {
			fmt.Printf("任务 %s: 成功, 执行时间: %v\n",
				result.ID, result.Duration)
		}
	}
}
