package main

// 设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。

import (
	"fmt"
	"sync"
	"time"
)

// Task 任务结构体，包含任务名称和执行函数
type Task struct {
	Name string
	Func func()
}

// TaskResult 任务执行结果
type TaskResult struct {
	TaskName   string
	Duration   time.Duration
	FinishTime time.Time
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks   []Task
	results chan TaskResult
	wg      sync.WaitGroup
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make(chan TaskResult, 10),
	}
}

// AddTask 添加任务到调度器
func (ts *TaskScheduler) AddTask(name string, taskFunc func()) {
	ts.tasks = append(ts.tasks, Task{Name: name, Func: taskFunc})
}

// Execute 执行所有任务并返回结果
func (ts *TaskScheduler) Execute() []TaskResult {
	results := make([]TaskResult, 0, len(ts.tasks))
	
	// 启动结果收集协程
	go func() {
		for result := range ts.results {
			results = append(results, result)
		}
	}()
	
	// 并发执行所有任务
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.executeTask(task)
	}
	
	// 等待所有任务完成
	ts.wg.Wait()
	close(ts.results)
	
	return results
}

// executeTask 执行单个任务并记录时间
func (ts *TaskScheduler) executeTask(task Task) {
	defer ts.wg.Done()
	
	start := time.Now()
	task.Func()
	duration := time.Since(start)
	
	ts.results <- TaskResult{
		TaskName:   task.Name,
		Duration:   duration,
		FinishTime: start,
	}
}

// 原有的奇数和偶数任务函数
func odd(num int) {
	for i := 1; i <= num; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 模拟任务耗时
	}
}

func even(num int) {
	for i := 2; i <= num; i += 2 {
		fmt.Printf("偶数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 模拟任务耗时
	}
}

func main() {
	// 创建任务调度器
	scheduler := NewTaskScheduler()
	
	// 添加任务
	scheduler.AddTask("打印奇数", func() { odd(10) })
	scheduler.AddTask("打印偶数", func() { even(10) })
	
	// 执行任务并获取结果
	fmt.Println("开始执行任务...")
	results := scheduler.Execute()
	
	// 输出执行结果
	fmt.Println("\n任务执行统计:")
	for _, result := range results {
		fmt.Printf("任务 '%s' 执行时间: %v\n", result.TaskName, result.Duration)
	}
}